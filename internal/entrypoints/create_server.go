package entrypoints

import (
	"inforce-test-task-service/internal/config"
	"inforce-test-task-service/internal/routes"

	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
)

func NewServer(dependencies *config.AppDependencies) *gin.Engine {
	server := gin.New()

	server.Use(sloggin.New(dependencies.Logger))
	server.Use(gin.Recovery())
	createServerRoutes(server, dependencies)

	return server
}

func createServerRoutes(server *gin.Engine, dependencies *config.AppDependencies) {
	routes.CreateRaribleRoutes(server, dependencies)
}