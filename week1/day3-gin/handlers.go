package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getProducts(c *gin.Context) {
	products := []gin.H{
		{"id": 1, "name": "Laptop", "price": 999},
		{"id": 2, "name": "Phone", "price": 699},
	}
	c.JSON(http.StatusOK, products)
}

func getProduct(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id": id, "name": "Product Name"})
}

func createProduct(c *gin.Context) {
	var product gin.H
	c.BindJSON(&product)
	c.JSON(http.StatusCreated, product)
}
