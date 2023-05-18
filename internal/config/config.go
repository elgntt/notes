package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PgUser     string
	PgPassword string
	PgHost     string
	PgPort     uint16
	PgDatabase string
	PgSSLMode  string

	LogFilePath string
	ServerPort  string
}

func GetConfig() (Config, error) {
	PgPort, err := strconv.ParseInt(getKey("PGPORT"), 0, 16)
	if err != nil {
		return Config{}, err
	}

	return Config{
		PgUser:     getKey("PGUSER"),
		PgPassword: getKey("PGPASSWORD"),
		PgHost:     getKey("PGHOST"),
		PgPort:     uint16(PgPort),
		PgDatabase: getKey("PGDATABASE"),
		PgSSLMode:  getKey("PGSSLMODE"),

		ServerPort:  getKey("PORT"),
		LogFilePath: getKey("LOG_FILE_PATH"),
	}, nil
}

func getKey(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
		return ""
	}

	return os.Getenv(key)
}
