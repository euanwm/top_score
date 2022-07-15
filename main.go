package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	const csvFp string = "381.csv"
	PrintTopTotal(csvFp)
}

func PrintTopTotal(resultFp string) {
	file, err := os.Open(resultFp)
	if err != nil {
		fmt.Println(err)
	}
	csvReader := csv.NewReader(file)
	results, _ := csvReader.ReadAll()
	for _, line := range results {
		fmt.Println(line[3], line[13])
	}
}
