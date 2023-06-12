package model


type Users struct {
	Username string `bson:"Username"`
	Password string `bson:"Password"`
}
type Laptops struct {
	Brand string `bson:"Brand"`
	Model string `bson:"Model"`
	Year  string `bson:"Year"`
	Price string `bson:"Price"`
}