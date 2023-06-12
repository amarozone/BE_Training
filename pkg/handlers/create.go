package handlers

import (
	"CRUD_API/pkg/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/myapp/services"
	"CRUD_API/pkg/services"
)

type LaptopHandler struct {
	service *services.LaptopService
}

func NewLaptopHandler(service *services.LaptopService) *LaptopHandler {
	fmt.Println("In NewLaptopHandler")
	return &LaptopHandler{
		service: service,
	}
}

func (h *LaptopHandler) PostLaptops(c *gin.Context) {
	fmt.Println("hello iam in PostLaptops handler")
	laptop := new(model.Laptops)
	if err := c.BindJSON(&laptop); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := h.service.CreateLaptop(laptop)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Posted successfully"})
}































































// package routes

// import (
// 	// getcollection "CRUD_API/Collection"
// 	// database "CRUD_API/databases"
// 	model "CRUD_API/model"
// 	"context"
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// )

// func PostLaptops(c *gin.Context) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	laptop := new(model.Laptops)
// 	defer cancel()
	
// 	if err := c.BindJSON(&laptop); err != nil {
// 	c.JSON(http.StatusBadRequest, gin.H{"message": err})
// 	log.Fatal(err)
// 	return
// 	}
// 	laptopPayload := model.Laptops{
// 		Brand: laptop.Brand,
// 		Model: laptop.Model,
// 		Year: laptop.Year,
// 		Price: laptop.Price,
// 	}
	
// 	err := s.store.PostLaptops(ctx, laptopPayload)
	
// 	if err != nil {
// 	c.JSON(http.StatusInternalServerError, gin.H{"message": err})
// 	return
// 	}
	
// 	c.JSON(http.StatusCreated, gin.H{"message": "Posted successfully", "Data": map[string]interface{}{"data": result}})
//    }