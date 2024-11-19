package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sampleCmd)
}

var sampleCmd = &cobra.Command{
	Use:   "sample",
	Short: "Generate a sample profile source in JSON",
	Long:  "Generate a sample profile source with complete information in JSON",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generating sample ...")
	},
}
