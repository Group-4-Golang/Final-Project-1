package controller

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func InitHtttpRoute(route *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}
	todosGroup := route.Group("/api/v1")

	todosGroup.GET("/todos", h.GetTodosHandler)
	todosGroup.GET("/todos/:id", h.GetTodoHandler)
	todosGroup.POST("/todos", h.CreateTodoHandler)
	todosGroup.PUT("/todos/:id", h.UpdateTodoHandler)
	todosGroup.DELETE("/todos/:id", h.DeleteTodoHandler)
	todosGroup.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
