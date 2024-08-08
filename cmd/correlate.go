package cmd

import (
	"github.com/ianhecker/correlate/internal/csv"
	"github.com/ianhecker/correlate/internal/oracle"
	"github.com/spf13/cobra"
)

var correlateCmd = &cobra.Command{
	Use:   "correlate",
	Short: "Create statistics from with CSV files",
	Run: func(cmd *cobra.Command, args []string) {
		correlate(InputFile, OutputFile)
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
