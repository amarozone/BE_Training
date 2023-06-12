package services



func (s *LaptopService) DeleteLaptop(laptopBrand string) error {
	err := s.store.DeleteLaptop(laptopBrand)
	if err != nil {
		// Handle any error or perform additional operations if needed
		return err
	}
	return nil
}
