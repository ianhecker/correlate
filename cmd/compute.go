package cmd

import (
	"github.com/ianhecker/correlate/internal/csv"
	"github.com/ianhecker/correlate/internal/oracle"
	"github.com/spf13/cobra"
)

var computeCmd = &cobra.Command{
	Use:   "compute",
	Short: "Compute statistics for an oracle",
}

func init() {
	rootCmd.AddCommand(computeCmd)
}

func correlate(oracle oracle.Oracle, inputFile string, outputFile string) {
	data, err := csv.ReadFile(inputFile)
	checkErr(err)

	txns, err := oracle.ParseMatrixIntoTransactions(data)
	checkErr(err)

	stats := txns.MakeStatistics()

	err = csv.WriteFile(outputFile, stats)
	checkErr(err)
}
