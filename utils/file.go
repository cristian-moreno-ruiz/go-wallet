package utils

import (
	"encoding/csv"
	"os"
)

// OpenCSV Opens the CSV
func OpenCSV(path string) [][]string {
	pwd, _ := os.Getwd()
	csvFile, _ := os.Open(pwd + path)
	defer csvFile.Close()
	csvLines, _ := csv.NewReader(csvFile).ReadAll()
	return csvLines
}
