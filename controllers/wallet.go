package controllers

import (
	"encoding/csv"
	"fmt"
	"net/http"

	"github.com/cristian-moreno-ruiz/go-wallet/services"
)

// Import File
func Import(w http.ResponseWriter, r *http.Request) {
	file, _, _ := r.FormFile("file")
	fmt.Println("hit", file)

	csvLines, _ := csv.NewReader(file).ReadAll()
	fmt.Println("file", csvLines)

	action := r.Form.Get("action")

	switch action {
	case "profit":
		buyOperations, sellOperations := services.ParseOperations(csvLines)

		sellOperations = services.CalculateProfit(buyOperations, sellOperations)
		fmt.Fprintln(w, sellOperations)
	case "save":
		buyOperations, sellOperations := services.ParseOperations(csvLines)
		fmt.Fprintln(w, "saving", buyOperations, sellOperations)

		for _, op := range buyOperations {
			op.Create()
		}
	}
}
