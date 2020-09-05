package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() Config {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}

	c := Config{}
	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into config struct, %v", err)
	}

	return c
}
