package handlers

import (
	"CRUD_API/pkg/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)



func (h *LaptopHandler) UpdateLaptop(c *gin.Context) {
	fmt.Println("in update laptop handler")
	laptopBrand := c.Param("laptopBrand")
	laptop := new(model.Laptops)
	if err := c.BindJSON(&laptop); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	err := h.service.UpdateLaptop(laptopBrand, laptop)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Data updated successfully!"})
}
