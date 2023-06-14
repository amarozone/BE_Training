package handlers

import (
	"CRUD_API/pkg/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetLaptops godoc
// @Summary      Get laptops by brand
// @Description  Retrieve a list of laptops by brand
// @Tags         laptops
// @Accept       json
// @Produce      json
// @Param        laptop body model.Laptops true "Laptop"
// @Router       /laptops/:laptopBrand [get]
func (h *LaptopHandler) GetLaptops(c *gin.Context) {
	laptop := new(model.Laptops)
	fmt.Println(laptop)
	laptopBrand := c.Param("laptopBrand")
	result, err := h.service.GetLaptop(laptopBrand)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "success!", "Data": result})
}
