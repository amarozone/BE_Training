package routes_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	routes "CRUD_API/routes"
)

func TestPostLaptops(t *testing.T) {
	// Setup test router
	router := gin.Default()
	router.POST("/laptops", routes.PostLaptops)

	// Prepare test data
	laptop := map[string]interface{}{
		"Brand": "Example Brand",
		"Model": "Example Model",
		"Year":  "2023",
		"Price": "199999",
	}
	payload, _ := json.Marshal(laptop)

	// Perform test request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/laptops", bytes.NewBuffer(payload))
	router.ServeHTTP(w, req)

	// Assert response
	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "Posted successfully", response["message"])

	data := response["Data"].(map[string]interface{})["data"]
	assert.NotNil(t, data)

	// Additional assertions based on expected data
	// assert.Equal(t, "Example Brand", data["Brand"])
	// assert.Equal(t, "Example Model", data["Model"])
	// assert.Equal(t, 2023, data["Year"])
	// assert.Equal(t, 1999.99, data["Price"])
}
