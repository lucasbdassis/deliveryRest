package routes

import (
	"delivery/controllers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		store := cookie.NewStore([]byte("secret"))
		store.Options(sessions.Options{MaxAge:   1}) // expire in a day

		product := main.Group("product")
		product.Use(sessions.Sessions("mysession", store))
		product.Use(controllers.LoginGet)

		{
			product.GET("/:id", controllers.GetProduct)
			product.GET("/", controllers.GetAllProducts)
			product.POST("/", controllers.CreateProduct)
			product.PUT("/", controllers.UpdateProduct)
			product.DELETE("/:id", controllers.DeleteProduct)

		}
		login := main.Group("")
		login.Use(sessions.Sessions("mysession", store))

		{
			login.POST("/login", controllers.LoginPost)
			login.GET("login/checked",controllers.LoginGet)
			login.POST("/register", controllers.Register)
		}

	}
	return router
}
