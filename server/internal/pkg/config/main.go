package config

import "github.com/spf13/viper"

var ApiAddress string
var PostgresUrl string

func init() {
	viper.SetDefault("api_address", ":4000")
	viper.SetDefault("postgres_url", "postgres:///kafka-chat")
	viper.AutomaticEnv()

	ApiAddress = viper.GetString("api_address")
	PostgresUrl = viper.GetString("postgres_url")
}
