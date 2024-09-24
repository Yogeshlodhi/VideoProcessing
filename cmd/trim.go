package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var trimCmd = &cobra.Command{
	Use: "trim [inputFile] [startTime] [endTime] [outputFile]",
	Short: "Trim video from start time to end time",
	Args: cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		inputFile := args[0]
		startTime := args[1]
		endTime := args[2]
		outputFile := args[3]
		err := trimVideo(inputFile, startTime, endTime, outputFile)
		if err != nil {
			fmt.Printf("Failed to trim it: %v\n", err)
		}
	},
}

func init(){
	rootCmd.AddCommand(trimCmd)
}

func trimVideo(input, start, end, output string) error {
	cmd := exec.Command("fmmpeg", "-i", input, "-ss", start, "-to", end, "-c", "copy", output)
	return cmd.Run()
}