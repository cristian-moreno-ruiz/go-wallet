package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Operations collection
var Operations *mongo.Collection

// BuyOperation describes a buy operation
type BuyOperation struct {
	Date         string
	BtcAmount    float64
	BtcPrice     float64
	FiatAmount   float64
	OperationFee float64
	BtcSold      float64
	ID           primitive.ObjectID `bson:"_id"`
}

// SellOperation describes a sell operation
type SellOperation struct {
	Date         string
	BtcAmount    float64
	BtcPrice     float64
	FiatAmount   float64
	OperationFee float64
	FiatBuyCost  float64
	FiatBuyFee   float64
	Profit       float64
}

func init() {
	Operations = database.Collection("operations")
	fmt.Println("Initializing Operations collection")
}

// Create Stores the operation into the Database
func (operation BuyOperation) Create() {
	fmt.Println(operation)
	Operations.InsertOne(ctx, operation)
}

// func (this SellOperation) create() {

// }
