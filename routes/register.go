package routes

import (
	// getcollection "CRUD_API/Collection"
	// database "CRUD_API/databases"
	// model "CRUD_API/model"
	// "context"
	// "log"
	// "net/http"
	// "time"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context){
	// var DB = database.ConnectDB()
	// var postCollection = getcollection.GetCollectionUser(DB, "Users")
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// user := new(model.Users)
	// defer cancel()

	// if err := c.BindJSON(&user); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"message": err})
	// 	log.Fatal(err)
	// 	return
	// }

	// newUser := model.Users{
	// 	Username: user.Username,
	// 	Password: user.Password,
	// }
    // result, err := postCollection.InsertOne(ctx, newUser)
    // if err != nil {
    //     c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register"})
    //     return
    // }

    // c.JSON(http.StatusCreated, gin.H{"message": "Posted successfully", "Data": map[string]interface{}{"data": result}})

}