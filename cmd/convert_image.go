package cmd

import (
	"fmt"
	"os/exec"
	"github.com/spf13/cobra"
)

var convertImageCmd = &cobra.Command{
	Use:   "convertImage [inputFile] [outputFile]",
	Short: "Convert an image to a different format",
	Args:  cobra.ExactArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		inputFile := args[0]
		outputFile := args[1]

		if err := convertImage(inputFile, outputFile); err != nil {
			fmt.Printf("Failed to convert image: %v\n", err)
		} else {
			fmt.Println("Image converted successfully!")
		}
	},
}

func init() {
	rootCmd.AddCommand(convertImageCmd)
}

func convertImage(input, output string) error {
	cmd := exec.Command("ffmpeg", "-i", input, output)
	return cmd.Run()
}