package services

import "CRUD_API/pkg/model"


func (s *LaptopService) GetLaptop(laptopBrand string) (*model.Laptops, error) {
	result, err := s.store.GetLaptop(laptopBrand)
	if err != nil {
		// Handle any error or perform additional operations if needed
		return nil, err
	}
	return result, nil
}
