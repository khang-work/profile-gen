package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/khang-work/profile-gen/sample"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sampleCmd)
	sampleCmd.Flags().StringP("output", "o", "", "Output file name")
}

var sampleCmd = &cobra.Command{
	Use:   "sample",
	Short: "Generate a sample profile source in JSON",
	Long:  "Generate a sample profile source with complete information in JSON",
	Run: func(cmd *cobra.Command, args []string) {
		outputFile, _ := cmd.Flags().GetString("output")

		if err := generateSampleProfile(outputFile); err != nil {
			log.Fatalf("Error generating sample profile: %s", err)
		}
	},
}

func generateSampleProfile(outputFile string) (retErr error) {
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

	// Marshal the profile into JSON.
	jsonData, err := json.MarshalIndent(profile, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal profile: %w", err)
	}

	// Write the JSON data to stdout (or any io.Writer).
	_, err = writer.Write(jsonData)
	if err != nil {
		return fmt.Errorf("write JSON to writer: %v", err)
	}

	return nil
}
