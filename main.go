package main

import (
	"ArticleSV/controller"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	address := fmt.Sprintf(":%s", port)
	route := gin.Default()
	route.Use(corsMiddleware())
	route.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "hello everynyan 😼",
		})
	})
	route.POST("/article", controller.CreateArticle)
	route.GET("/article", controller.RetrieveArticles)
	route.GET("/article/:id", controller.RetrieveArticle)
	route.PUT("/article/:id", controller.UpdateArticle)
	route.DELETE("/article/:id", controller.RemoveArticle)
	route.Run(address)
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		} else {
			c.Next()
		}
	}
}
