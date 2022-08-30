package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var Env ENV

func InitDotEnv() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	appPort := GetEnvString("APP_PORT")
	database := Database{URI: GetEnvString("DB_URI"), Name: GetEnvString("DB_Name")}

	Env = ENV{
		AppPort:  appPort,
		Database: database,
	}
}

func GetEnvString(key string) string {
	return os.Getenv(key)
}

func GetEnv() *ENV {
	return &Env
}
