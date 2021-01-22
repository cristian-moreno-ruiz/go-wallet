package services

import (
	"fmt"
)

// BuyOperation describes a buy operation
type BuyOperation struct {
	Date         string
	BtcAmount    float64
	BtcPrice     float64
	FiatAmount   float64
	OperationFee float64
	BtcSold      float64
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

// CalculateProfit Calculates the profit of a passed in combination off buys and sells
func CalculateProfit(buyOperations []BuyOperation, sellOperations []SellOperation) []SellOperation {
	// fmt.Println("       Date          BTC Amt   BTC Price   Fiat Sell  Fee   FIFO Cost   FIFO FEE   Net Profit")
	for i := range sellOperations {
		currentSellOp := &sellOperations[i]
		fmt.Println(sellOperations[i].Date)

		// Step 1. Find the buy cost, following FIFO policy
		bought := 0.0
		buyOp := 0

		// TODO: Add safety here, in case we entered more sells than buys by mistake, to not loop forever
		for bought < currentSellOp.BtcAmount {
			btcRemaining := buyOperations[buyOp].BtcAmount - buyOperations[buyOp].BtcSold
			if btcRemaining > 0 {
				// Still not sold, let's use it

				// If there is enough to use this buy operation for the whole sell amount, use it all and finish
				if btcRemaining >= currentSellOp.BtcAmount-bought {

					// If we don´t use an slice of pointers, we should access to the actual value, not its copy coming from the range operator
					// sellOperations[i].fiatBuyCost = currentSellOp.btcAmount * buyOperations[buyOp].btcPrice
					currentSellOp.FiatBuyCost = currentSellOp.FiatBuyCost + (currentSellOp.BtcAmount-bought)*buyOperations[buyOp].BtcPrice

					// Compute Fee
					currentSellOp.FiatBuyFee = currentSellOp.FiatBuyFee + (currentSellOp.BtcAmount-bought)*buyOperations[buyOp].OperationFee/buyOperations[buyOp].BtcAmount

					// Mark the amount used sold and consider its buying done
					bought = bought + currentSellOp.BtcAmount - bought
					buyOperations[buyOp].BtcSold = buyOperations[buyOp].BtcSold + currentSellOp.BtcAmount - bought
				} else {
					// Maybe should use Max[currentSellOp.fiatBuyCost, btcRemaining*buyOperations[buyOp].btcPrice]
					// In this way, it would be in my favor tax-wise
					currentSellOp.FiatBuyCost = currentSellOp.FiatBuyCost + btcRemaining*buyOperations[buyOp].BtcPrice

					// This should mark the buy operation as completely sold
					buyOperations[buyOp].BtcSold = buyOperations[buyOp].BtcSold + btcRemaining

					// Compute Fee
					currentSellOp.FiatBuyFee = currentSellOp.FiatBuyFee + buyOperations[buyOp].OperationFee

					bought = bought + btcRemaining
				}
			}
			buyOp++
		}
		currentSellOp.Profit = currentSellOp.FiatAmount - currentSellOp.FiatBuyCost - currentSellOp.FiatBuyFee - currentSellOp.OperationFee
	}

	fmt.Println("finished")
	fmt.Println(sellOperations)
	return sellOperations
}
