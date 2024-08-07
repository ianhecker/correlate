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

	// for i := 0; i < len(data); i++ {
	// 	for j := 0; j < len(data[i]); j++ {
	// 		fmt.Printf("i:%d j:%d thing:%s\n", i, j, data[i][j])
	// 	}
	// }

	// fmt.Println(data)

	txns, err := oracle.ParseCSV(data)
	checkErr(err)

	for k, v := range txns {
		fmt.Printf("index: %d data: %+v\n", k, v)
	}
}

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
