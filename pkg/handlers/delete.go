package handlers

import (
	"CRUD_API/pkg/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteLaptop godoc
// @Summary      Delete a laptop
// @Description  Delete a laptop by brand
// @Tags         laptops
// @Accept       json
// @Produce      json
// @Param        laptop body model.Laptops true "Laptop"
// @Router       /laptop/:laptopBrand [delete]
func (h *LaptopHandler) DeleteLaptop(c *gin.Context) {
	laptop := new(model.Laptops)
	fmt.Println(laptop)
	laptopBrand := c.Param("laptopBrand")
	err := h.service.DeleteLaptop(laptopBrand)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Article deleted successfully"})
}
