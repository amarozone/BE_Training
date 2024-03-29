package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var DB *mongo.Database

func GetDB() *mongo.Database {
	if DB ==nil {
		DB = ConnectDB()
	}
	return DB
}
func ConnectDB() *mongo.Database {

 Mongo_URL := "mongodb://root:password@127.0.0.1:27017"

 client, err := mongo.NewClient(options.Client().ApplyURI(Mongo_URL))
 
 if err != nil {
 log.Fatal(err)
 }
 
 ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
 err = client.Connect(ctx)
 defer cancel()
 
 if err != nil {
 log.Fatal(err)
 }
 
 fmt.Println("Connected to mongoDB")
 return client.Database("laptopDB")
}



































// package database

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )
 
// func ConnectDB() *mongo.Client {

//  Mongo_URL := "mongodb://root:password@127.0.0.1:27017"

//  client, err := mongo.NewClient(options.Client().ApplyURI(Mongo_URL))
 
//  if err != nil {
//  log.Fatal(err)
//  }
 
//  ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
//  err = client.Connect(ctx)
//  defer cancel()
 
//  if err != nil {
//  log.Fatal(err)
//  }
 
//  fmt.Println("Connected to mongoDB")
//  return client
// }





