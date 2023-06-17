package main

import (
	_ "finalproject1/docs"
	"finalproject1/router"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// @title Final Project 1 [Go + Gin Todo API]
// @version 1.0
// @description This is a simple server todo server.

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi
func main() {
	r := gin.Default()
	v1 := r.Group("/api/v1")

	router.TodorRoutes(v1)
	v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
