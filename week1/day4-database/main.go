package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	InitDB()

	r := gin.Default()

	r.GET("/products", getProducts)

	r.POST("/products", createProduct)

	r.Run(":8080")
}
