package config

import (
	"time"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	DB_USERNAME string `env:"DB_USERNAME,required"`
	DB_PASSWORD string `env:"DB_PASSWORD,required"`
	DB_NAME     string `env:"DB_NAME,required"`
	DB_HOST     string `env:"DB_HOST,required"`
	DB_PORT     string `env:"DB_PORT,required"`
	ORIGIN      string `env:"ORIGIN,required"`
	ENV         string `env:"ENV,required"`
	Port        string `env:"PORT,required"`

	AccessTokenPrivateKey  string        `env:"ACCESS_TOKEN_PRIVATE_KEY"`
	RefreshTokenPrivateKey string        `env:"REFRESH_TOKEN_PRIVATE_KEY"`
	AccessTokenExpiresIn   time.Duration `env:"ACCESS_TOKEN_EXPIRED_IN"`
	RefreshTokenExpiresIn  time.Duration `env:"REFRESH_TOKEN_EXPIRED_IN"`
	AccessTokenMaxAge      int           `env:"ACCESS_TOKEN_MAXAGE"`
	RefreshTokenMaxAge     int           `env:"REFRESH_TOKEN_MAXAGE"`

	GoogleClientID         string `env:"GOOGLE_OAUTH_CLIENT_ID"`
	GoogleClientSecret     string `env:"GOOGLE_OAUTH_CLIENT_SECRET"`
	GoogleOAuthRedirectUrl string `env:"GOOGLE_OAUTH_REDIRECT_URL"`
}

func LoadConfig() (config Config, err error) {
	err = godotenv.Load(".env.local")
	if err != nil {
		err := godotenv.Load(".env")
		if err != nil {
			panic("Error loading .env file")
		}
	}
	godotenv.Load()
	err = env.Parse(&config)

	return
}
