package stores

import (
	database "CRUD_API/databases"
	"CRUD_API/pkg/model"
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type LaptopStore struct {
	DB *mongo.Database
}

func NewLaptopStore() *LaptopStore {
	fmt.Println("In NewLaptopStore")
	return &LaptopStore{
		DB:database.GetDB(),
	}
}

func (s *LaptopStore) PostLaptops(laptop *model.Laptops) error {
	postCollection := s.DB.Collection("Laptops")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	laptopPayload := model.Laptops{
		Brand: laptop.Brand,
		Model: laptop.Model,
		Year:  laptop.Year,
		Price: laptop.Price,
	}

	_, err := postCollection.InsertOne(ctx, laptopPayload)
	if err != nil {
		// Handle any error or perform additional operations if needed
		return err
	}
	return nil
}

func (s *LaptopStore) UpdateLaptop(laptopBrand string, laptop *model.Laptops) error {
	postCollection := s.DB.Collection("Laptops")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	edited := bson.M{"Brand": laptop.Brand, "Model": laptop.Model, "Year": laptop.Year, "Price": laptop.Price}

	result, err := postCollection.UpdateOne(ctx, bson.M{"Brand": laptopBrand}, bson.M{"$set": edited})
	if err != nil {
		// Handle any error or perform additional operations if needed
		return err
	}

	if result.MatchedCount < 1 {
		// Handle case where data doesn't exist
		return errors.New("Data doesn't exist")
	}

	return nil
}

func (s *LaptopStore) DeleteLaptop(laptopBrand string) error {
	postCollection := s.DB.Collection("Laptops")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := postCollection.DeleteOne(ctx, bson.M{"Brand": laptopBrand})
	if err != nil {
		// Handle any error or perform additional operations if needed
		return err
	}

	if result.DeletedCount < 1 {
		// Handle case where no data is deleted
		return errors.New("No data to delete")
	}

	return nil
}

func (s *LaptopStore) GetLaptop(laptopBrand string) (*model.Laptops, error) {
	postCollection := s.DB.Collection("Laptops")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var result model.Laptops
	err := postCollection.FindOne(ctx, bson.M{"Brand": laptopBrand}).Decode(&result)
	if err != nil {
		// Handle any error or perform additional operations if needed
		return nil, err
	}

	return &result, nil
}





























































// package stores

// import (
// 	database "CRUD_API/databases"
// 	"CRUD_API/pkg/model"
// 	"context"

// 	"go.mongodb.org/mongo-driver/mongo"
// )

// type store struct {
// 	DB *mongo.Database
// }



// type Store interface {
// 	PostLaptops(laptop model.Laptops) error
// 	UpdateLaptop(initialLaptopBrand, updatedLaptopBrand model.Laptops) error
// 	ListOneLaptop(laptopBrand int) (error, model.Laptops)
// 	DeleteLaptop(laptopBrand int) error
// }

// func NewStore() Store {
// 	return &store{
// 		DB: database.GetDB(),
// 	}
// }

// func (s *store) PostLaptops(ctx context.Context, laptop model.Laptops) error {
// 	var postCollection = s.DB.Collection("Laptops")
// 	 postCollection.InsertOne(ctx, laptop)

// }

// func (s *store) UpdateLaptop(initialMovie, updatedMovie model.MovieTable) error {
// 	return s.DB.Model(&initialMovie).Updates(updatedMovie).Error
// }

// func (s *store) DeleteLaptop(id int) error {
// 	return s.DB.Delete(&model.MovieTable{}, id).Error
// }

// func (s *store) ListOneLaptop(id int) (error, model.MovieTable) {
// 	var movie model.MovieTable
// 	return s.DB.First(&movie, id).Error, movie

// }
