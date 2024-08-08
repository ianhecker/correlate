package cmd

import (
	"github.com/ianhecker/correlate/internal/csv"
	"github.com/ianhecker/correlate/internal/oracle"
	"github.com/spf13/cobra"
)

var correlateCmd = &cobra.Command{
	Use:   "correlate [input CSV] [output CSV]",
	Short: "Create statistics with CSV files",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		inputFile := args[0]
		outputFile := args[1]
		correlate(inputFile, outputFile)
	},
}

func init() {
	rootCmd.AddCommand(correlateCmd)
}

func correlate(inputFile string, outputFile string) {
	data, err := csv.ReadFile(inputFile)
	checkErr(err)

	txns, err := oracle.ParseMatrixIntoTransactions(data)
	checkErr(err)

	stats := txns.Statistics()

	err = csv.WriteFile(outputFile, stats)
	checkErr(err)
}
