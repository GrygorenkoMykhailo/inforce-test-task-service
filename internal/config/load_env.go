package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(logger *slog.Logger) {
	envFilePath := os.Getenv("APP_ENV_FILE_PATH")

	if envFilePath == "" {
		envFilePath = "./.env"
	}

	err := godotenv.Load(envFilePath)
	if err != nil {
		logger.Error("Failed to parse env", "err", err)
		os.Exit(1)
	}
}