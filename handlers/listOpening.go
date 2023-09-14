package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET opening",
	})
}
