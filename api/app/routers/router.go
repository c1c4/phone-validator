package routers

import (
	"api/app/controllers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.GET("/phones", controllers.GetAllPhones)
	}
}
