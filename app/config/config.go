package config

import (
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	PgDb       string
	PgUser     string
	PgPassword string
	PgHost     string
	PgPort     string

	JWTSecret            string
	JWTExpirationSeconds int
}

func InitConfig() (*config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &config{
		PgDb:                 os.Getenv("POSTGRES_DB"),
		PgUser:               os.Getenv("POSTGRES_USER"),
		PgPassword:           os.Getenv("POSTGRES_PASSWORD"),
		PgHost:               os.Getenv("POSTGRES_HOST"),
		PgPort:               os.Getenv("POSTGRES_PORT"),
		JWTSecret:            os.Getenv("JWT_SECRET"),
		JWTExpirationSeconds: 60 * 60 * 24 * 7,
	}, nil
}

var Config, _ = InitConfig()
