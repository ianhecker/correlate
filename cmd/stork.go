package cmd

import (
	"github.com/ianhecker/correlate/internal/oracle/stork"
	"github.com/spf13/cobra"
)

var computeStorkCmd = &cobra.Command{
	Use:   "stork",
	Short: "Compute statistics for Stork",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		inputFile := args[0]
		outputFile := args[1]
		correlate(stork.Stork{}, inputFile, outputFile)
	},
}

func init() {
	computeCmd.AddCommand(computeStorkCmd)
}
