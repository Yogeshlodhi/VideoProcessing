package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "videoprocess",
	Short: "This is a modern tool to play with videos",
	Long: "A CLI tool built in golang to alter/edit, change video options",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the VideoProcessing Tool!")
	},
}

func Execute(){
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}