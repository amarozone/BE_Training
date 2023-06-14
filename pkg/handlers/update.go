package handlers

import (
	"CRUD_API/pkg/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


// UpdateLaptop godoc
// @Summary      Update a laptop by brand
// @Description  Update the details of a laptop by brand
// @Tags         laptops
// @Accept       json
// @Produce      json
// @Param        laptop body model.Laptops true "Laptop"
// @Router       /laptops/:laptopBrand [put]
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
