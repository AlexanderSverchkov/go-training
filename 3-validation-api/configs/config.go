package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SmtpConfig
}

type SmtpConfig struct {
	SmtpEmail    string
	SmtpPassword string
	SmtpAddress  string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	return &Config{
		SmtpConfig: SmtpConfig{
			SmtpEmail:    os.Getenv("SMTP_EMAIL"),
			SmtpPassword: os.Getenv("SMTP_PASSWORD"),
			SmtpAddress:  os.Getenv("SMTP_ADDRESS"),
		},
	}
}
