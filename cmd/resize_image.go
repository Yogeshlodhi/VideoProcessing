package cmd

import (
	"fmt"
	"os/exec"
	"github.com/spf13/cobra"
)

var resizeImageCmd = &cobra.Command{
	Use:   "resizeImage [inputFile] [outputFile] [width] [height]",
	Short: "Resize an image to the specified width and height",
	Args:  cobra.ExactArgs(4),

	Run: func(cmd *cobra.Command, args []string) {
		inputFile := args[0]
		outputFile := args[1]
		width := args[2]
		height := args[3]

		if err := resizeImage(inputFile, outputFile, width, height); err != nil {
			fmt.Printf("Failed to resize image: %v\n", err)
		} else {
			fmt.Println("Image resized successfully!")
		}
	},
}

func init() {
	rootCmd.AddCommand(resizeImageCmd)
}

func resizeImage(input, output, width, height string) error {
	cmd := exec.Command("ffmpeg", "-i", input, "-vf", fmt.Sprintf("scale=%s:%s", width, height), output)
	return cmd.Run()
}
