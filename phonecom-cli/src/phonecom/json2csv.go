package main

import (
	"fmt"
	"github.com/yukithm/json2csv"
	"os"
)

func exportToCsv(outputJson interface{}) {

	results, err := json2csv.JSON2CSV(outputJson)

	if err != nil {
		fmt.Println("Error while converting JSON to CSV")
	}
	if len(results) == 0 {
		return
	}

	csv := json2csv.NewCSVWriter(os.Stdout)
	csv.HeaderStyle = json2csv.DotNotationStyle
	csv.Transpose = false

	if err := csv.WriteCSV(results); err != nil {
		fmt.Println("Error while writing result to CSV")
	}
}
