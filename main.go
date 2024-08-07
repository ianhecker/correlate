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

func main() {
	data, err := readCSVFile("./chainlink.csv")
	checkErr(err)

	txns, err := oracle.ParseCSV(data)
	checkErr(err)

	file, err := os.Create("out.csv")
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(file)
	defer w.Flush()

	rows := [][]string{
		{"ID", "Mean", "Median", "Max", "Min", "Standard Deviation"},
		txns.TimeStatsToStrings(),
		txns.C1StatsToStrings(),
		txns.C2StatsToStrings(),
		txns.CostOfC1InUSDStatsToStrings(),
		txns.CostOfC2InUSDStatsToStrings(),
	}

	for _, row := range rows {
		if err := w.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
}

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
