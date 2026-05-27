package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Status": "OK",
		})
	})

	r.GET("/products", getProducts)
	r.GET("/products/:id", getProduct)
	r.POST("/products", createProduct)

	r.Run(":8080")
}
