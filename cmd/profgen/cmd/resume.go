package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/khang-work/profile-gen/resumedr/txt"
	"github.com/khang-work/profile-gen/sample"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(resumeCmd)
	resumeCmd.Flags().StringP("format", "f", "txt", "Resume output format")
	resumeCmd.Flags().StringP("output", "o", "", "Resume output file name")
}

var resumeCmd = &cobra.Command{
	Use:   "resume",
	Short: "Generate the resume in selected format",
	Long:  "Generate the resume in one of the available formats (txt)",
	Run: func(cmd *cobra.Command, args []string) {
		outputFormat, _ := cmd.Flags().GetString("format")
		outputFile, _ := cmd.Flags().GetString("output")

		if err := generateResume(outputFile, outputFormat); err != nil {
			log.Fatalf("Error generating resume: %s", err)
		}
	},
}

func generateResume(outputFile string, outputFormat string) (retErr error) {
	writer := os.Stdout

	// Use file instead of stdout if file name is provided.
	if outputFile != "" {
		file, err := os.Create(outputFile)
		if err != nil {
			return fmt.Errorf("create file: %w", err)
		}

		defer func() {
			if err := file.Close(); err != nil {
				retErr = fmt.Errorf("close file: %w", err)
			}
		}()

		writer = file
	}
	profile := sample.GenerateSample()

	switch outputFormat {
	case "txt":
		resumeBlob := txt.ProfileToPlainTextResume(&profile)
		fmt.Fprintf(writer, resumeBlob)
	}

	return nil
}
