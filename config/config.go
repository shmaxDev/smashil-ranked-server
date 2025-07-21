package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvVars struct {
	DbConnectionString string
	Port string
	DbType string
}

func GetEnv() (EnvVars) {
	err := godotenv.Load()

	if err != nil{
		log.Fatal("error loading .env file")
	}

	return EnvVars{os.Getenv("DB_CONNECTION_STRING"), os.Getenv("PORT"), os.Getenv("DB_TYPE")}
}