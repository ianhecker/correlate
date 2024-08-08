package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/ianhecker/correlate/internal/oracle"
)

func readCSVFile(filepath string) ([][]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("Could not find file: %s", filepath)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	data, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("Unable to parse file: %s: %w", filepath, err)
	}

	return data, nil
}

func writeCSVFile(filepath string, matrix [][]string) error {
	file, err := os.Create(filepath)
	defer file.Close()

	if err != nil {
		log.Fatalln("Failed to open file", err)
	}

	w := csv.NewWriter(file)
	defer w.Flush()

	for _, row := range matrix {
		if err := w.Write(row); err != nil {
			return fmt.Errorf("error writing record to file: %w", err)
		}
	}
	return nil
}

func main() {
	data, err := readCSVFile("./chainlink.csv")
	checkErr(err)

	txns, err := oracle.ParseCSV(data)
	checkErr(err)

	stats := txns.Statistics()

	err = writeCSVFile("./chainlink-out.csv", stats)
	checkErr(err)
}

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
