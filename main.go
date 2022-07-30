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
		writeCSV(lifterName, lifterHistory)
	}
}

func bigListPrint(listBoi [][]string) {
	for _, contents := range listBoi {
		fmt.Println(contents)
	}
}

//genFilename Getting wavy and writing code, blank return cuz i'm edgy and quirky
func genFilename(lifternName string) (filename string) {
	filename = strings.ReplaceAll(lifternName, " ", "_") + ".csv"
	return
}

//writeCSV Writes CSV file, first arg is the filepath/name. Second is the bigSlice data.
func writeCSV(lifterName string, bigSlice [][]string) {
	csvName := genFilename(lifterName)
	newCsvFile, err := os.Create(csvName)
	if err != nil {
		log.Fatal(err)
	}
	writer := csv.NewWriter(newCsvFile)
	writeData := writer.WriteAll(bigSlice)
	if writeData != nil {
		fmt.Println(writeData)
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
