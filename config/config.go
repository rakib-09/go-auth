package config

import (
	"github.com/spf13/viper"
	"log"
)

type AppConfig struct {
	Name string
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

type Config struct {
	App *AppConfig
	DB  *DBConfig
}

var config *Config

func DB() *DBConfig {
	return config.DB
}

func App() *AppConfig {
	return config.App
}

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	viper.Unmarshal(&config)

	config.App = &AppConfig{
		Name: viper.GetString("APP_NAME"),
		Port: viper.GetString("APP_URL"),
	}

	config.DB = &DBConfig{
		Host:     viper.GetString("DB_HOST"),
		Port:     viper.GetString("DB_PORT"),
		User:     viper.GetString("DB_USER"),
		Pass:     viper.GetString("DB_PASSWORD"),
		Database: viper.GetString("DB_DATABASE"),
	}
}
