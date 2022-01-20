// Package spreadsheet provides functions for using spreadsheets
package spreadsheet

import (
	"encoding/csv"
	"log"
	"os"
)

// ReadSheet provides a method for reading spreadsheet data from a csv file
func ReadSheet(name string) [][]string {
	f, err := os.Open(name)

	if err != nil {
		log.Fatalf("Cannot open '%s': %s\n", name, err.Error())
	}

	defer f.Close()

	r := csv.NewReader(f)

	r.Comma = ','

	rows, err := r.ReadAll()

	if err != nil {
		log.Fatalln("Cannot read CSV data:", err.Error())
	}

	return rows

}
