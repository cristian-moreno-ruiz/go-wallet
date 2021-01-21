package controllers

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type buyOperation struct {
	date         string
	btcAmount    float64
	btcPrice     float64
	fiatAmount   float64
	operationFee float64
	btcSold      float64
}

type sellOperation struct {
	date         string
	btcAmount    float64
	btcPrice     float64
	fiatAmount   float64
	operationFee float64
	fiatBuyCost  float64
	fiatBuyFee   float64
	profit       float64
}

var buyOperations []buyOperation
var sellOperations []sellOperation

// Open s the file
func Open( /* w http.ResponseWriter, r *http.Request */ ) {
	pwd, _ := os.Getwd()
	csvFile, _ := os.Open(pwd + "/test/raw_operations.csv")
	csvLines, _ := csv.NewReader(csvFile).ReadAll()

	buyOperations = make([]buyOperation, 0, len(csvLines)-1)
	buyOperations = make([]buyOperation, 0, len(csvLines)-1)
	fmt.Println(buyOperations)
	for _, v := range csvLines[1:] {

		if v[0] != "" {
			btcAmount, _ := strconv.ParseFloat(v[1], 64)
			btcPrice, _ := strconv.ParseFloat(v[2], 64)
			fiatAmount, _ := strconv.ParseFloat(v[3], 64)
			operationFee, _ := strconv.ParseFloat(v[4], 64)
			buyOperations = append(buyOperations, buyOperation{v[0], btcAmount, btcPrice, fiatAmount, operationFee, 0})
		}

		if v[5] != "" {
			btcAmount, _ := strconv.ParseFloat(v[6], 64)
			btcPrice, _ := strconv.ParseFloat(v[7], 64)
			fiatAmount, _ := strconv.ParseFloat(v[8], 64)
			operationFee, _ := strconv.ParseFloat(v[9], 64)
			sellOperations = append(sellOperations, sellOperation{v[5], btcAmount, btcPrice, fiatAmount, operationFee, 0, 0, 0})
		}
	}
	// fmt.Println(buyOperations)
	// fmt.Println(sellOperations)
}

// CalculateProfit Calculate the profit obtained taking into account only the BTC that have already been sold
func CalculateProfit() {
	for i, currentSellOp := range sellOperations {
		fmt.Println(sellOperations[i].date)

		// Step 1. Find the buy cost, following FIFO policy
		bought := 0.0
		buyOp := 0

		for bought < currentSellOp.btcAmount {
			btcRemaining := buyOperations[buyOp].btcAmount - buyOperations[buyOp].btcSold
			if btcRemaining > 0 {
				// Still not sold, let's use it

				// If there is enough to use this buy operation for the whole sell amount, use it all and finish
				if btcRemaining >= currentSellOp.btcAmount {
					bought = bought + currentSellOp.btcAmount
					sellOperations[i].fiatBuyCost = currentSellOp.btcAmount * buyOperations[buyOp].btcPrice
					// currentSellOp.fiatBuyCost = currentSellOp.btcAmount * buyOperations[buyOp].btcPrice
					fmt.Println(currentSellOp)
				}
			}
		}
	}

	fmt.Println(sellOperations)
}
