package routers

import (
	"AsiaYo_BackendEngineer/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	router := gin.Default()
	router.GET("/convert_currency", controllers.ConvertCurrency)

	return router
}
