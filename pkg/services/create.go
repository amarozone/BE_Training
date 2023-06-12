package services

import (
	"CRUD_API/pkg/model"
	"fmt"

	"CRUD_API/pkg/stores"
)

type LaptopService struct {
	store *stores.LaptopStore
}

func NewLaptopService(store *stores.LaptopStore) *LaptopService {
	fmt.Println("In NewLaptopService")
	return &LaptopService{
		store: store,
	}
}

func (s *LaptopService) CreateLaptop(laptop *model.Laptops) error {
	err := s.store.PostLaptops(laptop)
	if err != nil {
		// Handle any error or perform additional operations if needed
		return err
	}
	return nil
}



























































// package services

// import (
// 	"CRUD_API/pkg/model"
// 	"context"
// 	"net/http"

// 	"github.com/pkg/errors"
// 	// "go.mongodb.org/mongo-driver/bson"
// 	// "go.mongodb.org/mongo-driver/bson/primitive"
// )

// func (s *service) PostLaptops(ctx context.Context, laptop model.Laptops) {
	
// 	err := s.store.PostLaptops(ctx, laptop)
// 	if err != nil {
// 		return model.Laptops{}, serror.NewServiceError(http.StatusInternalServerError, errors.Wrap(err, "failed to create filter"))
// 	}
// 	return nil

// }
