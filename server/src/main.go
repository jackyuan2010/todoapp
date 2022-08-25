package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/cors"
)

func index(c *gin.Context) {
	c.Header("Access-Control-Allow", "*")
	c.JSON(http.StatusOK, gin.H{
			"message": "TODO App",
		})
}

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/", index)

	return router
}

func main() {
	fmt.Println("to do app starting....")
	router := SetupRoutes()
	router.Run(":8081")
}