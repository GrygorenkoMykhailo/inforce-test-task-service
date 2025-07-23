package routes

import (
	"inforce-test-task-service/internal/config"
	"inforce-test-task-service/internal/controllers"

	"github.com/gin-gonic/gin"
)

func CreateRaribleRoutes(server *gin.Engine, dependencies *config.AppDependencies) {
	group := server.Group("/rarible")

	controller := controllers.NewRaribleController(dependencies)
	group.GET("/nft-ownership/:collection_id", controller.GetNftOwnershipsById)
	group.POST("/nft-rarity", controller.GetNftTraitsRarity)
}