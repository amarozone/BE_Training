package handlers

import (

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"CRUD_API/model"
	"CRUD_API/pkg/handlers"
	"CRUD_API/pkg/services"
	"CRUD_API/pkg/stores"
)

func TestPostLaptops(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// db := database.ConnectDB()
	laptopStore := stores.NewLaptopStore()
	laptopService := services.NewLaptopService(laptopStore)
	laptopHandler := handlers.NewLaptopHandler(laptopService)
	// Set up the route handler
	router.POST("/laptop", laptopHandler.PostLaptops)

	// Create a test laptop
	laptop := model.Laptops{
		Brand: "Test Brand",
		Model: "Test Model",
		Year:  "2023",
		Price: "999",
	}

	// Convert the laptop to JSON
	laptopJSON, _ := json.Marshal(laptop)

	// Create a request with the laptop JSON as the body
	// w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/laptop", strings.NewReader(string(laptopJSON)))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to capture the response
	rec := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(rec, req)

	// Assert the status code
	assert.Equal(t, http.StatusCreated, rec.Code)

	// Assert the response body
	expectedResponse := `{"message":"Posted successfully"}`
	assert.JSONEq(t, expectedResponse, rec.Body.String())
}




func TestGetLaptops(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// db := database.ConnectDB()
	laptopStore := stores.NewLaptopStore()
	laptopService := services.NewLaptopService(laptopStore)
	laptopHandler := handlers.NewLaptopHandler(laptopService)
	// Define the GET route
	router.GET("/laptop/:laptopBrand", laptopHandler.GetLaptops)

	// Create a new HTTP request for the GET route with a valid laptop brand
	req, err := http.NewRequest("GET", "/laptop/Test Brand", nil)
	assert.NoError(t, err)

	// Create a new HTTP response recorder
	res := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(res, req)

	// Assert the response status code
	assert.Equal(t, http.StatusCreated, res.Code)

	// Assert the response body
	expected := `{"Data":{"Brand":"Test Brand","Model":"Test Model","Year":"2023","Price":"999"},"message":"success!"}`
	assert.Equal(t, expected, res.Body.String())

	// Create a new HTTP request for the GET route with a nonexistent laptop brand
	req, err = http.NewRequest("GET", "/laptop/nonexistentbrand", nil)
	assert.NoError(t, err)

	// Reset the response recorder
	res = httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(res, req)

	// Assert the response status code
	assert.Equal(t, http.StatusInternalServerError, res.Code)

	// Assert the response body
	expected = `{"message":"mongo: no documents in result"}`
	assert.Equal(t, expected, res.Body.String())
}




func TestUpdateLaptop(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// db := database.ConnectDB()
	laptopStore := stores.NewLaptopStore()
	laptopService := services.NewLaptopService(laptopStore)
	laptopHandler := handlers.NewLaptopHandler(laptopService)
	// Set up the route handler
	// Define the PUT route
	router.PUT("/laptop/:laptopBrand", laptopHandler.UpdateLaptop)

	// Create a new HTTP request for the PUT route with a valid laptop brand and request body
	reqBody := `{"Brand":"Updated Brand","Model":"Updated Model","Year":"2022","Price":"899"}`
	req, err := http.NewRequest("PUT", "/laptop/Test Brand", strings.NewReader(reqBody))
	assert.NoError(t, err)

	// Create a new HTTP response recorder
	res := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(res, req)

	// Assert the response status code
	assert.Equal(t, http.StatusCreated, res.Code)

	// Assert the response body
	expected := `{"message":"Data updated successfully!"}`
	assert.JSONEq(t, expected, res.Body.String())

	// Create a new HTTP request for the PUT route with a nonexistent laptop brand
	reqBody = `{"Brand":"Updated Brand","Model":"Updated Model","Year":"2022","Price":"899"}`
	req, err = http.NewRequest("PUT", "/laptop/nonexistentbrand", strings.NewReader(reqBody))
	assert.NoError(t, err)

	// Reset the response recorder
	res = httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(res, req)

	// Assert the response status code
	assert.Equal(t, http.StatusInternalServerError, res.Code)

	// Assert the response body
	expected = `{"message":"Data doesn't exist"}`
	assert.JSONEq(t, expected, res.Body.String())
}





func TestDeleteLaptop(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// db := database.ConnectDB()
	laptopStore := stores.NewLaptopStore()
	laptopService := services.NewLaptopService(laptopStore)
	laptopHandler := handlers.NewLaptopHandler(laptopService)
	// Set up the route handler
	// Define the DELETE route
	router.DELETE("/laptop/:laptopBrand", laptopHandler.DeleteLaptop)

	// Create a new HTTP request for the DELETE route
	req, err := http.NewRequest("DELETE", "/laptop/Updated Brand", nil)
	assert.NoError(t, err)

	// Create a new HTTP response recorder
	res := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(res, req)

	// Assert the response status code
	assert.Equal(t, http.StatusCreated, res.Code)

	// Assert the response body
	expected := `{"message":"Article deleted successfully"}`
	
	assert.Equal(t, expected, res.Body.String())


	// Create a new HTTP request for the DELETE route with a nonexistent laptop brand
	req, err = http.NewRequest("DELETE", "/laptop/nonexistentbrand", nil)
	assert.NoError(t, err)

	// Reset the response recorder
	res = httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(res, req)

	// Assert the response status code
	assert.Equal(t, http.StatusInternalServerError, res.Code)

	// Assert the response body
	expected = `{"message":"No data to delete"}`
	assert.Equal(t, expected, res.Body.String())
	
}
