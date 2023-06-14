package main

import (
	// database "CRUD_API/databases"
	handlers "CRUD_API/pkg/handlers"
	"CRUD_API/pkg/services"
	"CRUD_API/pkg/stores"
	routes "CRUD_API/routes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	// "github.com/swaggo/gin-swagger"
)

// gin-swagger middleware
// swagger embed files
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the JWT token from the request header

		tokenString1 := c.GetHeader("Authorization")
		tokenString := tokenString1[7:]
		fmt.Println("this is middleware token", tokenString)
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

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	router := gin.Default()

	// db := database.ConnectDB()
	laptopStore := stores.NewLaptopStore()
	laptopService := services.NewLaptopService(laptopStore)
	laptopHandler := handlers.NewLaptopHandler(laptopService)
	router.POST("/laptop", laptopHandler.PostLaptops)
	router.PUT("laptop/:laptopBrand", laptopHandler.UpdateLaptop)
	router.DELETE("laptop/:laptopBrand", laptopHandler.DeleteLaptop)
	router.GET("laptop/:laptopBrand", laptopHandler.GetLaptops)

	// public.POST("/register", routes.Register)
	// public.POST("/login",routes.Login)
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
