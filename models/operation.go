package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

// Operations collection
var Operations *mongo.Collection

const (
	// SELL constant
	SELL = "SELL"
	// BUY constant
	BUY = "BUY"
)

// BuyOperation describes a buy operation
type BuyOperation struct {
	Date         string
	BtcAmount    float64
	BtcPrice     float64
	FiatAmount   float64
	OperationFee float64
	BtcSold      float64
	Type         string
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
	Type         string
}

func init() {
	Operations = database.Collection("operations")
	fmt.Println("Initializing Operations collection")
}

/*
BUY OPERATIONS
*/

// Create Stores the Buy operation into the Database
func (operation BuyOperation) Create() {
	operation.Type = BUY
	Operations.InsertOne(ctx, &operation)
}

/*
SELL OPERATIONS
*/

// Create Stores the Sell operation into the Database
func (operation SellOperation) Create() {
	operation.Type = SELL
	Operations.InsertOne(ctx, operation)
}
