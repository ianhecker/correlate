package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/ianhecker/correlate/internal/correlate"
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

	// for i := 0; i < len(data); i++ {
	// 	for j := 0; j < len(data[i]); j++ {
	// 		fmt.Printf("i:%d j:%d thing:%s\n", i, j, data[i][j])
	// 	}
	// }

	// fmt.Println(data)

	txns, err := oracle.ParseCSV(data)
	checkErr(err)

	fmt.Printf("txns.ID: %+v\n", txns.ID)
	fmt.Printf("txns.Date: %+v\n", txns.Date)
	fmt.Printf("txns.Time: %+v\n", txns.Time)
	fmt.Printf("txns.C1: %+v\n", txns.C1)
	fmt.Printf("txns.C2: %+v\n", txns.C2)
	fmt.Printf("txns.C1_USD: %+v\n", txns.C1_USD)
	fmt.Printf("txns.C2_USD: %+v\n", txns.C2_USD)

	fmt.Printf("txns.C1InUSD: %+v\n", txns.CostOfC1InUSD())
	fmt.Printf("txns.C2InUSD: %+v\n", txns.CostOfC2InUSD())

	c1StdDev := correlate.StandardDeviation(txns.C1...)
	c2StdDev := correlate.StandardDeviation(txns.C2...)
	c1USDStdDev := correlate.StandardDeviation(txns.CostOfC1InUSD()...)
	c2USDStdDev := correlate.StandardDeviation(txns.CostOfC2InUSD()...)

	fmt.Printf("c1 std dev: %f\n", c1StdDev)
	fmt.Printf("c2 std dev: %f\n", c2StdDev)
	fmt.Printf("c1 USD std dev: %f\n", c1USDStdDev)
	fmt.Printf("c2 USD std dev: %f\n", c2USDStdDev)
}

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
