/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"exemplar-api/internal/config"
	"exemplar-api/internal/server"
	"fmt"
	"log"
	"net/http"

	migrate "github.com/rubenv/sql-migrate"
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

	// initialise our live-reload config
	config.InitConfig()

	testPrint := viper.GetString("testPrint")

	port := viper.GetString("port")

	log.Println("Test Print: ", testPrint)

	// get migrations
	migrations := &migrate.FileMigrationSource{
		Dir: "./migrations",
	}

	// conbnect to DB
	// Connection string
	connStr := "host=localhost port=5432 user=myuser password=mypassword dbname=mydb sslmode=disable"

	// Open the connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// Ping to test connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	fmt.Println("Successfully connected to PostgreSQL!")

	// perform migrations
	log.Println("Performing migrations!")

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Println("error performing migrations: ", err)
	}
	log.Printf("Applied %d migrations", n)

	handler := server.NewHandler(db)

	listenStr := fmt.Sprintf(":%s", port)

	log.Println("API listening on", listenStr)
	if err = http.ListenAndServe(listenStr, handler); err != nil {
		log.Panic(err)
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
