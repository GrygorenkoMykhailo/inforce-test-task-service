package config

import (
	"inforce-test-task-service/internal/apis"
	"log/slog"
)

type AppDependencies struct {
	Logger *slog.Logger
	RaribleApi apis.RaribleApi
}