package main

import (
	"inforce-test-task-service/internal/config"
)

func main() {
	logger := config.NewLogger()
	config.LoadEnv(logger)
}