package routes

import (
	"github.com/Carmo-sousa/api/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/api/v1")
	{
		users := main.Group("users")
		{
			users.GET("/", controllers.ShowUsers)
			users.POST("/", controllers.AddUser)
			users.GET("/:name", controllers.ShowUser)
			users.DELETE("/:email", controllers.DeleteUser)
		}
	}

	return router
}
