package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"joelovien/computer-based-test/routes"
)

func main() {
	port := os.Getenv("PORT") //check for the declared port in the env

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	//Register our routes
	routes.AuthRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": "Access granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": "Access granted for api-2"})
	})

	router.Run(":" + port)
}
