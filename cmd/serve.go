/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"database/sql"
	"exemplar-api/internal/config"
	"exemplar-api/internal/migrations"
	"exemplar-api/internal/server"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/amacneil/dbmate/v2/pkg/dbmate"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: serve,
}

func serve(cmd *cobra.Command, args []string) {

	// make a channel for signalling reloaded config
	reloadCh := make(chan struct{})
	cfg := &config.Config{}

	cfg.InitConfig(reloadCh)

	// conbnect to DB
	// Connection string

	for {

		// parse DSN to url
		dsnUrl, err := buildDatabaseURL("postgres", cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBUser, cfg.DBPassword, cfg.DBSSLMode)
		if err != nil {
			log.Fatalln(err)
		}

		u, err := url.Parse(dsnUrl.String())
		if err != nil {
			log.Println("error parsing dsnurl: ", err)
		}
		dbMigrations := dbmate.New(u)

		// get migrations
		dbMigrations.FS = migrations.EmbeddedMigrations

		migrations, err := dbMigrations.FindMigrations()
		if err != nil {
			log.Println("error finding migrations:", err)
		}
		for _, m := range migrations {
			log.Println(m.Version, m.FilePath)
		}

		log.Println("Applying Migrations...")

		err = dbMigrations.CreateAndMigrate()
		if err != nil {
			panic(err)
		}

		// Open the connection
		db, err := sql.Open("postgres", cfg.DSN)
		if err != nil {
			log.Fatalf("Failed to open database: %v", err)
		}

		// Ping to test connection
		if err := db.Ping(); err != nil {
			log.Fatalf("Failed to ping database: %v", err)
		}

		r := http.NewServeMux()

		handler := server.NewHandler(db)
		r.HandleFunc("GET /notes", handler.ListNotes)
		r.HandleFunc("POST /notes", handler.CreateNote)
		r.HandleFunc("GET /notes", handler.GetNote)
		r.HandleFunc("PUT /notes", handler.CreateNote)
		r.HandleFunc("DELETE /notes", handler.CreateNote)

		listenStr := fmt.Sprintf(":%s", cfg.Port)
		title := cfg.Title

		srv := &http.Server{
			Addr:    listenStr,
			Handler: r,
		}

		go func() {

			log.Println("Title: ", title)

			log.Println("API listening on", listenStr)
			if err = srv.ListenAndServe(); err != nil {
				// dont panic here, or you'll kill the whole application due to no recover()!
				log.Println(err)
			}
		}()

		<-reloadCh

		log.Println("Config has been changed, reloading...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		_ = srv.Shutdown(ctx)
		cancel()

		// reload config
		err = viper.Unmarshal(cfg)
		if err != nil {
			log.Println(err)
		}
	}

}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Function to build Database URL
func buildDatabaseURL(driver, host, port, dbname, user, password, sslmode string) (*url.URL, error) {
	if host == "" || port == "" || dbname == "" || user == "" || password == "" {
		return nil, fmt.Errorf("missing required database connection parameters")
	}

	u := &url.URL{
		Scheme: driver,
		User:   url.UserPassword(user, password),
		Host:   fmt.Sprintf("%s:%s", host, port),
		Path:   "/" + dbname,
	}

	if sslmode != "" {
		query := url.Values{}
		query.Add("sslmode", sslmode)
		u.RawQuery = query.Encode()
	}

	return u, nil
}
