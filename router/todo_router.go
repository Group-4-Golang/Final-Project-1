package router

import (
	"finalproject1/controller"

	"github.com/gin-gonic/gin"
)

func TodorRoutes(route *gin.RouterGroup) {
	route.GET("/todos", controller.GetTodosHandler)
	route.GET("/todos/:id", controller.GetTodoHandler)
	route.POST("/todos", controller.CreateTodoHandler)
	route.PUT("/todos/:id", controller.UpdateTodoHandler)
	route.DELETE("/todos/:id", controller.DeleteTodoHandler)
}
