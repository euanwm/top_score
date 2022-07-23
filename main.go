package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

func main() {
	var lifterName = strings.Join(os.Args[1:], " ")
	const csvDir string = "UK"
	csvFiles, err := os.ReadDir(csvDir)
	if err != nil {
		log.Fatal(err)
	}
	var lifterHistory [][]string
	for _, csvPath := range csvFiles {
		var result = GetLifterResult(path.Join(csvDir, csvPath.Name()), lifterName)
		if len(result) != 0 {
			lifterHistory = append(lifterHistory, result)
		}
	}
	if len(lifterHistory) == 0 {
		fmt.Println("No results found!")
	} else {
		bigListPrint(lifterHistory)
	}
}

func bigListPrint(listBoi [][]string) {
	for _, contents := range listBoi {
		fmt.Println(contents)
	}
}

func GetLifterResult(resultFp string, lifter string) []string {
	var lifterResult []string
	file, err := os.Open(resultFp)
	if err != nil {
		fmt.Println(err)
	}
	csvReader := csv.NewReader(file)
	results, _ := csvReader.ReadAll()
	for _, line := range results {
		if strings.ToLower(line[3]) == lifter {
			lifterResult = line
		}
	}
	return lifterResult
}
