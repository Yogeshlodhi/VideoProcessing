package cmd

import (
	"fmt"
	"os/exec"
	"github.com/spf13/cobra"
)

// watermarkImageCmd overlays a watermark onto an image
var watermarkImageCmd = &cobra.Command{
	// ================== DOES NOT WORK CURRENTLY =========================
	Use:   "watermarkImage [inputFile] [watermarkFile] [outputFile]",
	Short: "Add a watermark to an image",
	Args:  cobra.ExactArgs(3),

	Run: func(cmd *cobra.Command, args []string) {
		inputFile := args[0]
		watermarkFile := args[1]
		outputFile := args[2]

		if err := watermarkImage(inputFile, watermarkFile, outputFile); err != nil {
			fmt.Printf("Failed to add watermark: %v\n", err)
		} else {
			fmt.Println("Watermark added successfully!")
		}
	},
}

func init() {
	rootCmd.AddCommand(watermarkImageCmd)
}

func watermarkImage(input, watermark, output string) error {
	cmd := exec.Command("ffmpeg", "-i", input, "-i", watermark, "-filter_complex", "overlay", output)
	return cmd.Run()
}
