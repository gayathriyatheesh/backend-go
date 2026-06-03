package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	InitDB()

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	r.POST("/register", register)
	r.POST("/login", login)

	protected := r.Group("")
	protected.Use(AuthMiddleware())

	protected.GET("/products", getProducts)
	protected.POST("/products", createProduct)

	r.Run(":8080")
}
