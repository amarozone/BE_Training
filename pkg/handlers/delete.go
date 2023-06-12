package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)



func (h *LaptopHandler) DeleteLaptop(c *gin.Context) {
	laptopBrand := c.Param("laptopBrand")
	err := h.service.DeleteLaptop(laptopBrand)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Article deleted successfully"})
}
