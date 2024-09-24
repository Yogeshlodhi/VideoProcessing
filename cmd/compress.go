package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var compressCmd = &cobra.Command{
	Use: "Compress [inputFile] [outputFile]",
	Short: "Compress video to reduce the files size",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		inputFile := args[0]
		outputFile := args[1]
		err := compressVideo(inputFile, outputFile)
		if err != nil {
			fmt.Printf("Failed to compress video: %v\n", err)
		}
	},
}

func init(){
	rootCmd.AddCommand(compressCmd)
}

func compressVideo(input, output string) error {
	cmd := exec.Command("ffmpeg", "-i", input, "-vcodec", "libx264", "-crf", "28", output)
	return cmd.Run()
}

