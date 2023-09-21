package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// inicializa default
	router := gin.Default()

	// inicializar rotas
	InitRoutes(router)

	// roda
	router.Run()
}
