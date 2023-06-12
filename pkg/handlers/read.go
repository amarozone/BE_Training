package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func (h *LaptopHandler) GetLaptops(c *gin.Context) {
	laptopBrand := c.Param("laptopBrand")
	result, err := h.service.GetLaptop(laptopBrand)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "success!", "Data": result})
}
