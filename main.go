package main

import (
	routes "CRUD_API/routes"
	"fmt"
	// "fmt"
	"net/http"

	// "fmt"
	// "net/http"
	// "time"

	"github.com/gin-gonic/gin"
	// "github.com/golang-jwt/jwt"
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



func main() {
    router := gin.Default()

    public := router.Group("/api")

	public.POST("/register", routes.Register)
	public.POST("/login",routes.Login)
	//received the token from login

	protected := router.Group("/api/admin")
	protected.Use(authMiddleware())
	protected.POST("/laptops", routes.PostLaptops)
	protected.GET("laptops/:laptopBrand", routes.GetLaptops)
	protected.PUT("laptops/:laptopBrand", routes.UpdateLaptop)
	protected.DELETE("laptops/:laptopBrand", routes.DeleteLaptop)

	router.Run("localhost:3000")
}



// func main() {
//     router := gin.Default()
 
	// router.POST("/laptops", routes.PostLaptops)
	// router.GET("laptops/:laptopBrand", routes.GetLaptops)
	// router.PUT("laptops/:laptopBrand", routes.UpdateLaptop)
	// router.DELETE("laptops/:laptopBrand", routes.DeleteLaptop)
//     router.Run("localhost: 3000")
	
// }