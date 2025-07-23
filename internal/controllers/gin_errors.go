package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Status500(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message": "something went wrong",
	})
}