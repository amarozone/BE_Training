package getcollection

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("laptopDB").Collection("Laptops")
	return collection
   }
   func GetCollectionUser(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("laptopDB").Collection("Users")
	return collection
   }