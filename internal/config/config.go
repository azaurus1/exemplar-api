package config

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Port       string `mapstructure:"port"`
	Title      string `mapstructure:"title"`
	DSN        string `mapstructure:"dsn"`
	DBHost     string `mapstructure:"db_host"`
	DBPort     string `mapstructure:"db_port"`
	DBName     string `mapstructure:"db_name"`
	DBUser     string `mapstructure:"db_user"`
	DBPassword string `mapstructure:"db_password"`
	DBSSLMode  string `mapstructure:"db_sslmode"`
}

func (c *Config) InitConfig(onChange chan struct{}) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // will be SNAP_DATA at runtime

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	if err := viper.Unmarshal(&c); err != nil {
		log.Fatalf("unable to decode config into struct: %v", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
		onChange <- struct{}{}
	})
}
