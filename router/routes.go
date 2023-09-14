package router

import (
	"github.com/gin-gonic/gin"
	"main.go/handlers"
)

func InitRoutes(r *gin.Engine) {
	v1 := r.Group("/api/")
	{
		v1.GET("/opening", handlers.ShowOpeningHandler)
		v1.POST("/opening", handlers.CreateOpeningHandler)
		v1.DELETE("/opening", handlers.DeleteOpeningHandler)
		v1.PUT("/opening", handlers.UpdateOpeningHandler)
		v1.GET("/openings", handlers.ListOpeningHandler)
	}
}
