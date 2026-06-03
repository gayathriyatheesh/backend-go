package main

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var JWT_SECRET = "your-secret-key-change-this"

func HashPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed)
}

func VerifyPassword(hashedPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

func GenerateToken(userID uint, email string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenString, _ := token.SignedString([]byte(JWT_SECRET))
	return tokenString
}

func VerifyToken(tokenString string) (uint, error) {
	token, err := jwt.ParseWithClaims(
		tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(JWT_SECRET), nil
		})
	if err != nil {
		return 0, err
	}
	claims := token.Claims.(*jwt.MapClaims)
	userID := uint((*claims)["user_id"].(float64))
	return userID, nil
}
