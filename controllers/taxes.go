package controllers

import (
	"fmt"
	"net/http"

	"github.com/cristian-moreno-ruiz/go-wallet/services"
	"github.com/cristian-moreno-ruiz/go-wallet/utils"
)

// TODO: Taxes was never a proper name, this will be deprecated, and we will keep this to do it through cmd
// Specifying the path name

// Handle s the file
func Handle(w http.ResponseWriter, r *http.Request) {
	csvLines := utils.OpenCSV("/test/raw_operations.csv")

	buyOperations, sellOperations := services.ParseOperations(csvLines)

	_, sellOperations = services.CalculateProfit(buyOperations, sellOperations)
	fmt.Fprintln(w, sellOperations)
}
