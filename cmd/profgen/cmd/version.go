package cmd

import (
	"fmt"

	"github.com/khang-work/profile-gen/info"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of profile-gen",
	Long:  "Print the version number of profile-gen",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(info.SoftwareTitle)
	},
}
