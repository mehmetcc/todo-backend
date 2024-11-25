package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mehmetcc/todo-backend/internal/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	userGroup := router.Group("/users")
	{
		userGroup.GET("/", controllers.GetUsers)
		userGroup.POST("/", controllers.CreateUser)
		userGroup.GET("/:id", controllers.GetUser)
		userGroup.PUT("/:id", controllers.UpdateUser)
		userGroup.DELETE("/:id", controllers.DeleteUser)
	}

	taskGroup := router.Group("/tasks")
	{
		taskGroup.GET("/", controllers.GetTasks)
		taskGroup.POST("/", controllers.CreateTask)
		taskGroup.GET("/:id", controllers.GetTask)
		taskGroup.PUT("/:id", controllers.UpdateTask)
		taskGroup.DELETE("/:id", controllers.DeleteTask)
	}

	return router
}
