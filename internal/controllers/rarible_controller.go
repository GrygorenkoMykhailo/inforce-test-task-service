package controllers

import (
	"inforce-test-task-service/internal/apis"
	"inforce-test-task-service/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RaribleController interface {
	GetNftOwnershipsById(ctx *gin.Context)
	GetNftTraitsRarity(ctx *gin.Context)
}

type raribleControllerImpl struct {
	Dependencies *config.AppDependencies
}

func NewRaribleController(dependencies *config.AppDependencies) RaribleController {
	return &raribleControllerImpl{ Dependencies: dependencies }
}

func (controller *raribleControllerImpl) GetNftOwnershipsById(ctx *gin.Context) {
	logger := controller.Dependencies.Logger

	collectionId, exists := ctx.Params.Get("collection_id")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Collection id is a required path param",
		})
		return
	}
	
	data, err := controller.Dependencies.RaribleApi.GetNftOwnershipsById(collectionId)
	if err != nil {
		logger.Error("Failed to get nft ownersips", "err", err)
		Status500(ctx)
		return
	}

	ctx.JSON(http.StatusOK, data)
}


func (controller *raribleControllerImpl) GetNftTraitsRarity(ctx *gin.Context) {
	logger := controller.Dependencies.Logger

	var payload apis.GetNftTraitsRarityPayload
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		logger.Error("Invalid request body", "err", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
		})
		return
	}

	data, err := controller.Dependencies.RaribleApi.GetNftTraitsRarity(&payload)
	if err != nil {
		logger.Error("Failed to get nft traits rarity", "err", err)
		Status500(ctx)
		return
	}

	ctx.JSON(http.StatusOK, data)
}
