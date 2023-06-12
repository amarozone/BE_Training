package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func authMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Get the JWT token from the request header

        tokenString1 := c.GetHeader("Authorization")
		tokenString := tokenString1[7:]
		fmt.Println("this is middleware token",tokenString)
        // Parse the token
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            // Provide the secret key used for signing the token
            return []byte("your_secret_key"), nil
        })

        // Check if there was an error or the token is invalid
        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        // If the token is valid, you can retrieve the claims and perform further actions
        // For example, you can extract the user ID from the claims and store it in the request context
        claims, _ := token.Claims.(jwt.MapClaims)
        Username := claims["Username"].(string)
        c.Set("Username", Username)

        // Continue to the next handler
        c.Next()
    }
}