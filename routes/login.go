package routes

import (
	getcollection "CRUD_API/Collection"
	database "CRUD_API/databases"
	model "CRUD_API/model"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
)


func generateToken(Username string) (string, error) {
	// Create the claims for the JWT token
	claims := jwt.MapClaims{
		"Username": Username,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
	}

	// Create a new token with the claims and the signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	// Replace "your_secret_key" with your actual secret key
	tokenString, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		return "", err
	}
	fmt.Println("this is login token",(tokenString))
	return tokenString, nil
}

func Login(c *gin.Context) {
    var DB = database.ConnectDB()
    var userCollection = getcollection.GetCollectionUser(DB, "Users")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var credentials model.Users 

    if err := c.ShouldBindJSON(&credentials); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Find the user with the provided username
    filter := bson.M{"Username": credentials.Username}
    var user model.Users
    err := userCollection.FindOne(ctx, filter).Decode(&user)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    // Compare the provided password with the stored password
    if user.Password != credentials.Password {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    // Generate JWT token
    token, err := generateToken(user.Username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}
