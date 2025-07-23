package main

import (
	"inforce-test-task-service/internal/apis"
	"inforce-test-task-service/internal/config"
	"inforce-test-task-service/internal/entrypoints"
	"os"
)

func main() {
	logger := config.NewLogger()
	config.LoadEnv(logger)

	dependencies := config.AppDependencies{
		Logger: logger,
		RaribleApi: apis.NewRaribleApi(),
	}

	server := entrypoints.NewServer(&dependencies)
	err := server.Run("127.0.0.1:" + os.Getenv("PORT"))
	if err != nil {
		logger.Error("Failed to start server", "err", err.Error())
	}
}