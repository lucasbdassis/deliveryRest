package routes

import (
	"delivery/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		product := main.Group("product")
		{
			product.GET("/:id", controllers.GetProduct)
			product.GET("/", controllers.GetAllProducts)
			product.POST("/", controllers.CreateProduct)
			product.POST("/login", controllers.LoginPost)
			product.PUT("/", controllers.UpdateProduct)
			product.DELETE("/:id", controllers.DeleteProduct)

		}
	}
	return router
}
