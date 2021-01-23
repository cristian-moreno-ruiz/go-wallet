package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

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

// ListBuyOperations lists all operations or only buys or only sells if specified
func ListBuyOperations() []*BuyOperation {

	filter := bson.D{
		primitive.E{Key: "type", Value: BUY},
	}

	cur, err := Operations.Find(ctx, filter)

	var ops []*BuyOperation

	if err != nil {
		return ops
	}

	for cur.Next(ctx) {
		var op BuyOperation
		err := cur.Decode(&op)
		if err != nil {
			return ops
		}

		ops = append(ops, &op)
	}

	if err := cur.Err(); err != nil {
		return ops
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	if len(ops) == 0 {
		return ops
	}

	return ops
}

/*
SELL OPERATIONS
*/

// Create Stores the Sell operation into the Database
func (operation SellOperation) Create() {
	operation.Type = SELL
	Operations.InsertOne(ctx, operation)
}

// ListSellOperations lists all operations or only buys or only sells if specified
func ListSellOperations() []*SellOperation {

	filter := bson.D{
		primitive.E{Key: "type", Value: SELL},
	}

	cur, err := Operations.Find(ctx, filter)

	var ops []*SellOperation

	if err != nil {
		return ops
	}

	for cur.Next(ctx) {
		var op SellOperation
		err := cur.Decode(&op)
		if err != nil {
			return ops
		}

		ops = append(ops, &op)
	}

	if err := cur.Err(); err != nil {
		return ops
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	if len(ops) == 0 {
		return ops
	}

	return ops
}
