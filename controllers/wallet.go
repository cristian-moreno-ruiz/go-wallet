package controllers

import (
	"encoding/csv"
	"fmt"
	"net/http"

	"github.com/cristian-moreno-ruiz/go-wallet/models"

	"github.com/cristian-moreno-ruiz/go-wallet/services"
)

// ImportOperations File
func ImportOperations(w http.ResponseWriter, r *http.Request) {
	file, _, _ := r.FormFile("file")
	fmt.Println("hit", file)

	csvLines, _ := csv.NewReader(file).ReadAll()
	fmt.Println("file", csvLines)

	buyOperations, sellOperations := services.ParseOperations(csvLines)

	action := r.Form.Get("action")

	switch action {
	case "profit":
		sellOperations = services.CalculateProfit(buyOperations, sellOperations)
		fmt.Fprintln(w, sellOperations)
	case "save":
		sellOperations = services.CalculateProfit(buyOperations, sellOperations)
		fmt.Fprintln(w, "saving", buyOperations, sellOperations)

		for _, op := range buyOperations {
			op.Create()
		}

		for _, op := range sellOperations {
			op.Create()
		}
	}
}

// ListOperations Lists the operations in the wallet
func ListOperations(w http.ResponseWriter, r *http.Request) {

	opType, _ := r.URL.Query()["type"]

	if opType[0] == "buy" || opType[0] == "all" {
		ops := models.ListBuyOperations()
		for _, v := range ops {
			fmt.Fprintln(w, *v)
		}
	}
	if opType[0] == "sell" || opType[0] == "all" {
		ops := models.ListSellOperations()
		for _, v := range ops {
			fmt.Fprintln(w, *v)
		}
	}
}
