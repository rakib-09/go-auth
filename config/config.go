package config

import (
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"log"
	"time"
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

type JwtCustomClaims struct {
	Name string `json:"name"`
	UID  uint   `json:"uid"`
	jwt.StandardClaims
}

type JwtConfig struct {
	AccessTokenSecret  string
	RefreshTokenSecret string
	AccessTokenExpiry  time.Duration
	RefreshTokenExpiry time.Duration
	ContextKey         string
}

type Config struct {
	App *AppConfig
	DB  *DBConfig
	Jwt *JwtConfig
}

var config *Config

func DB() *DBConfig {
	return config.DB
}

func App() *AppConfig {
	return config.App
}

func Jwt() *JwtConfig {
	return config.Jwt
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

	config.Jwt = &JwtConfig{
		AccessTokenSecret:  "accesstokensecret",
		RefreshTokenSecret: "refreshtokensecret",
		AccessTokenExpiry:  300,
		RefreshTokenExpiry: 10080,
		ContextKey:         "user",
	}
}
