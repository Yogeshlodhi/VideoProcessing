package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var convertCmd = &cobra.Command{
	Use: "convert [inputFile] [outputFormat]",
	Short: "Convert video to another format like mp4, mkv",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		inputFile := args[0]
		outputFormat := args[1]

		fileBase := strings.TrimSuffix(inputFile, filepath.Ext(inputFile))
		// outputFile := fmt.Sprintf("%s.%s", inputFile, outputFormat)
		outputFile := fmt.Sprintf("%s.%s", fileBase, outputFormat)

		err := convertVideo(inputFile, outputFile)
		if err != nil {
			fmt.Printf("Failed to convert the video: %v\n", err)
		}else{
			fmt.Printf("Video successfully converted to %s\n", outputFile)
		}
	},
}

func init(){
	rootCmd.AddCommand(convertCmd)
}

func convertVideo(input, output string) error {
	cmd := exec.Command("ffmpeg", "-i", input, output)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}