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
