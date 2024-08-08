package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

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
