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

func register(c *gin.Context) {
	var req RegisterRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var existingUser User
	if err := DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		c.JSON(400, gin.H{"error": "Email already registered"})
		return
	}
	user := User{
		Name:     req.Name,
		Email:    req.Email,
		Password: HashPassword(req.Password),
	}
	DB.Create(&user)
	c.JSON(201, user)
}

func login(c *gin.Context) {
	var req LoginRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var user User
	if err := DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	if !VerifyPassword(user.Password, req.Password) {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	token := GenerateToken(user.ID, user.Email)
	c.JSON(200, LoginResponse{Token: token, User: user})
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		tokenString := authHeader[7:]
		userID, err := VerifyToken(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		c.Next()
	}
}
