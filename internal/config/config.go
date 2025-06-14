package config

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Port  string `mapstructure:"port"`
	Title string `mapstructure:"title"`
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
