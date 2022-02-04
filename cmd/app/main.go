package main

import (
	"learn/easyrest/pkg/controllers"
	"learn/easyrest/pkg/models"

	"github.com/gin-gonic/gin"
)

import "github.com/swaggo/gin-swagger" // gin-swagger middleware
import "github.com/swaggo/files"       // swagger embed files
import _ "learn/easyrest/docs"

// @title        Book API
// @version      0.0.1
// @description  Simple REST API for managing books

// @host      localhost:8080
// @BasePath  /

func main() {
	r := gin.Default()
	models.ConnectDB()

	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.PUT("/books/:id", controllers.UpdateBook)
	r.POST("/books", controllers.CreateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
