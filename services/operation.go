package services

import (
	"fmt"
	"strconv"
)

// ParseOperations Parses operations into buy and sell operations types
func ParseOperations(csvLines [][]string) ([]BuyOperation, []SellOperation) {
	buyOperations := make([]BuyOperation, 0, len(csvLines)-1)
	sellOperations := make([]SellOperation, 0, len(csvLines)-1)
	fmt.Println(buyOperations)
	for _, v := range csvLines[1:] {

		if v[0] != "" {
			btcAmount, _ := strconv.ParseFloat(v[1], 64)
			btcPrice, _ := strconv.ParseFloat(v[2], 64)
			fiatAmount, _ := strconv.ParseFloat(v[3], 64)
			operationFee, _ := strconv.ParseFloat(v[4], 64)
			buyOperations = append(buyOperations, BuyOperation{v[0], btcAmount, btcPrice, fiatAmount, operationFee, 0})
		}

		if v[5] != "" {
			btcAmount, _ := strconv.ParseFloat(v[6], 64)
			btcPrice, _ := strconv.ParseFloat(v[7], 64)
			fiatAmount, _ := strconv.ParseFloat(v[8], 64)
			operationFee, _ := strconv.ParseFloat(v[9], 64)
			sellOperations = append(sellOperations, SellOperation{v[5], btcAmount, btcPrice, fiatAmount, operationFee, 0, 0, 0})
		}
	}

	return buyOperations, sellOperations
}
