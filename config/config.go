package config

import "github.com/spf13/viper"

type config struct {
	PORT        string
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
}

var ENV *config

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		panic("Error reading config file: "+ err.Error())
	}

	if err := viper.Unmarshal(&ENV); err != nil {
		panic("Error unmarshalling config: " + err.Error())
	}
}