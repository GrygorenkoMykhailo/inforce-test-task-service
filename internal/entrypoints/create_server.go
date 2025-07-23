package entrypoints

import (
	"inforce-test-task-service/internal/config"
	"inforce-test-task-service/internal/routes"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
)

func NewServer(dependencies *config.AppDependencies) *gin.Engine {
	server := gin.New()

	server.Use(sloggin.New(dependencies.Logger))
	server.Use(gin.Recovery())
	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{ os.Getenv("FRONTEND_ORIGIN") },
		AllowMethods: []string{ "GET", "POST" },
		AllowCredentials: true,
	}))

	createServerRoutes(server, dependencies)

	return server
}

func createServerRoutes(server *gin.Engine, dependencies *config.AppDependencies) {
	routes.CreateRaribleRoutes(server, dependencies)
}