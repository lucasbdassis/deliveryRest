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
			product.PUT("/", controllers.UpdateProduct)
			product.DELETE("/:id", controllers.DeleteProduct)

		}
		login := main.Group("")
		{
			login.POST("/login", controllers.LoginPost)
			login.POST("/register", controllers.Register)
		}
	}
	return router
}
