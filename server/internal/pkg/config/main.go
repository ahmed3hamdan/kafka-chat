package config

import "github.com/spf13/viper"

var ApiAddress string
var PostgresUrl string
var JwtSecret string
var KafkaAddress string

func init() {
	viper.SetDefault("api_address", ":4000")
	viper.SetDefault("postgres_url", "postgres:///kafka-chat")
	viper.SetDefault("kafka_address", "kafka:9092")
	viper.AutomaticEnv()

	ApiAddress = viper.GetString("api_address")
	PostgresUrl = viper.GetString("postgres_url")
	JwtSecret = viper.GetString("jwt_secret")
	KafkaAddress = viper.GetString("kafka_address")
}
