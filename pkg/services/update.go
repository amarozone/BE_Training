package services

import (
	"CRUD_API/pkg/model"
	"fmt"
)


func (s *LaptopService) UpdateLaptop(laptopBrand string, laptop *model.Laptops) error {
	err := s.store.UpdateLaptop(laptopBrand, laptop)
	fmt.Println("in update laptop service")
	if err != nil {
		// Handle any error or perform additional operations if needed
		return err
	}
	return nil
}
