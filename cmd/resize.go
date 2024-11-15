package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var compressImageCmd = &cobra.Command{
	Use:   "compressImage [inputFile] [outputFile]",
	Short: "Compress image to reduce the file size",
	Args:  cobra.ExactArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		inputFile := args[0]
		outputFile := args[1]
		
		if err := compressImage(inputFile, outputFile); err != nil {
			fmt.Printf("Failed to compress image: %v\n", err)
		} else {
			fmt.Println("Image compressed successfully!")
		}
	},
}

func init() {
	rootCmd.AddCommand(compressImageCmd)
}


func compressImage(input, output string) error {
	cmd := exec.Command("ffmpeg", "-i", input, "-qscale:v", "5", output)
	return cmd.Run()
}