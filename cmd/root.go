package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var InputFile string
var OutputFile string

var rootCmd = &cobra.Command{
	Use:   "correlate",
	Short: "A brief description of your application",
}

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func Execute() {
	err := rootCmd.Execute()
	checkErr(err)
}

func init() {
	rootCmd.PersistentFlags().StringVar(&InputFile, "input", "in.csv", "Input CSV file")
	rootCmd.PersistentFlags().StringVar(&OutputFile, "output", "out.csv", "Output CSV file")
}
