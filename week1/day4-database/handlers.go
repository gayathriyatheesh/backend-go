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
	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash Password"})
		return
	}
	user := User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
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

	token, err := GenerateToken(user.ID, user.Email)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}
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

		if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			c.JSON(401, gin.H{"error": "Invalid authorization header"})
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

func addToCart(c *gin.Context) {
	userID, _ := c.Get("user_id")
	userIDuint := userID.(uint)
	var req AddToCartRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var product Product
	if err := DB.First(&product, req.ProductID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	var cart Cart
	DB.Where("user_id = ?", userIDuint).FirstOrCreate(&cart, Cart{UserID: userIDuint})

	var cartItem CartItem
	if err := DB.Where("cart_id = ? AND product_id = ?", cart.ID, req.ProductID).First(&cartItem).Error; err == nil {

		cartItem.Quantity += req.Quantity
		DB.Save(&cartItem)
	} else {

		cartItem := CartItem{
			CartID:    cart.ID,
			ProductID: req.ProductID,
			Quantity:  req.Quantity,
			Price:     product.Price,
		}
		DB.Create(&cartItem)
	}
	c.JSON(200, gin.H{"message": "Added to cart"})
}

func updateCartItem(c *gin.Context) {
	itemID := c.Param("id")
	var req struct {
		Quantity int `json:"quantity" binding:"required,min=1"`
	}
	c.BindJSON(&req)

	DB.Model(&CartItem{}).Where("id = ?", itemID).Update("quantity", req.Quantity)
	c.JSON(200, gin.H{"message": "Updated"})
}

func removeFromCart(c *gin.Context) {
	itemID := c.Param("id")
	DB.Delete(&CartItem{}, itemID)
	c.JSON(200, gin.H{"message": "Removed from cart"})
}
