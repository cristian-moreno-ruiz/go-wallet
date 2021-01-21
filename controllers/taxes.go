package controllers

import (
	"encoding/csv"
	"fmt"
	"net/http"
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
func Open(w http.ResponseWriter, r *http.Request) {
	pwd, _ := os.Getwd()
	//csvFile, _ := os.Open(pwd + "/test/raw_operations_complex.csv")
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

	CalculateProfit(w)
}

// CalculateProfit Calculate the profit obtained taking into account only the BTC that have already been sold
func CalculateProfit(w http.ResponseWriter) {
	fmt.Fprintln(w, "       Date          BTC Amt   BTC Price   Fiat Sell  Fee   FIFO Cost   FIFO FEE   Net Profit")
	for i := range sellOperations {
		currentSellOp := &sellOperations[i]
		fmt.Println(sellOperations[i].date)

		// Step 1. Find the buy cost, following FIFO policy
		bought := 0.0
		buyOp := 0

		// TODO: Add safety here, in case we entered more sells than buys by mistake, to not loop forever
		for bought < currentSellOp.btcAmount {
			btcRemaining := buyOperations[buyOp].btcAmount - buyOperations[buyOp].btcSold
			if btcRemaining > 0 {
				// Still not sold, let's use it

				// If there is enough to use this buy operation for the whole sell amount, use it all and finish
				if btcRemaining >= currentSellOp.btcAmount-bought {

					// If we don´t use an slice of pointers, we should access to the actual value, not its copy coming from the range operator
					// sellOperations[i].fiatBuyCost = currentSellOp.btcAmount * buyOperations[buyOp].btcPrice
					currentSellOp.fiatBuyCost = currentSellOp.fiatBuyCost + (currentSellOp.btcAmount-bought)*buyOperations[buyOp].btcPrice

					// Compute Fee
					currentSellOp.fiatBuyFee = currentSellOp.fiatBuyFee + (currentSellOp.btcAmount-bought)*buyOperations[buyOp].operationFee/buyOperations[buyOp].btcAmount

					// Mark the amount used sold and consider its buying done
					bought = bought + currentSellOp.btcAmount - bought
					buyOperations[buyOp].btcSold = buyOperations[buyOp].btcSold + currentSellOp.btcAmount - bought

					// fmt.Println(currentSellOp)
				} else {

					// Maybe should use Max[currentSellOp.fiatBuyCost, btcRemaining*buyOperations[buyOp].btcPrice]
					// In this way, it would be inmy favor tax-wise
					currentSellOp.fiatBuyCost = currentSellOp.fiatBuyCost + btcRemaining*buyOperations[buyOp].btcPrice

					// This should mark the buy operation as completely sold
					buyOperations[buyOp].btcSold = buyOperations[buyOp].btcSold + btcRemaining

					// Compute Fee
					currentSellOp.fiatBuyFee = currentSellOp.fiatBuyFee + buyOperations[buyOp].operationFee

					bought = bought + btcRemaining
					// fmt.Println(currentSellOp)
					// fmt.Println("ELSE")
				}
			}
			buyOp++
		}
		currentSellOp.profit = currentSellOp.fiatAmount - currentSellOp.fiatBuyCost - currentSellOp.fiatBuyFee - currentSellOp.operationFee
		fmt.Fprintln(w, *currentSellOp)
	}

	fmt.Println("finished")
	fmt.Println(sellOperations)
}
