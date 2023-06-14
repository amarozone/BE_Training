package handlers

import (
	"CRUD_API/pkg/model"
	"CRUD_API/pkg/stores"
	"testing"
	"github.com/stretchr/testify/assert"
)



func TestPostLaptop(t *testing.T) {
	// Create a new LaptopStore instance
	laptopStore := stores.NewLaptopStore()

	// Create a new laptop object
	laptop := &model.Laptops{
		Brand: "Test Brand",
		Model: "Test Model",
		Year:  "2023",
		Price: "999",
	}

	// Call the PostLaptops function
	err := laptopStore.PostLaptops(laptop)

	// Assert that no error occurred
	assert.NoError(t, err)
}

func TestUpdateLaptops(t *testing.T) {
	// Create a new LaptopStore instance
	laptopStore := stores.NewLaptopStore()

	// Define the laptop brand to update
	laptopBrand := "Test Brand"

	// Create a new laptop object with updated values
	updatedLaptop := &model.Laptops{
		Brand: "Updated Brand",
		Model: "Updated Model",
		Year:  "2022",
		Price: "899",
	}

	// Call the UpdateLaptop function
	err := laptopStore.UpdateLaptop(laptopBrand, updatedLaptop)

	// Assert that no error occurred
	assert.NoError(t, err)
}

func TestDeleteLaptops(t *testing.T) {
	// Create a new LaptopStore instance
	laptopStore := stores.NewLaptopStore()

	// Define the laptop brand to delete
	laptopBrand := "Test Brand"

	// Call the DeleteLaptop function
	err := laptopStore.DeleteLaptop(laptopBrand)

	// Assert that no error occurred
	assert.NoError(t, err)
}

func TestGetLaptop(t *testing.T) {
	// Create a new LaptopStore instance
	laptopStore :=stores.NewLaptopStore()

	// Define the laptop brand to retrieve
	laptopBrand := "Test Brand"

	// Call the GetLaptop function
	laptop, err := laptopStore.GetLaptop(laptopBrand)

	// Assert that no error occurred
	assert.NoError(t, err)

	// Assert that the retrieved laptop is not nil
	assert.NotNil(t, laptop)

	// Assert the retrieved laptop values
	expectedLaptop := &model.Laptops{
		Brand: "Test Brand",
		Model: "Test Model",
		Year:  "2023",
		Price: "999",
	}
	assert.EqualValues(t, expectedLaptop, laptop)}


