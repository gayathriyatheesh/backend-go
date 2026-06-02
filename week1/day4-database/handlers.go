package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getProducts(c *gin.Context) {

	var products []Product

	DB.Find(&products)

	c.JSON(http.StatusOK, products)
}

func createProduct(c *gin.Context) {

	var product Product

	c.BindJSON(&product)

	DB.Create(&product)

	c.JSON(http.StatusCreated, product)
}
