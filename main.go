package main

import (
	"finalproject1/configs"
	"finalproject1/controller"
	_ "finalproject1/docs"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	godotenv.Load()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		HttpMainHandler()
		defer wg.Done()
	}()
	wg.Wait()
}

func HttpMainHandler() {
	r := gin.Default()
	db := configs.NewConnection(configs.BaseConfig()).Database

	controller.InitHtttpRoute(r, db)

	r.Run()
}
