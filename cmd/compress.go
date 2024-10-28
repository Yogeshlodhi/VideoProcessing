package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)
// cobra.Command ( it is a struct ) represents a single command within a single CLI app
var compressCmd = &cobra.Command{
	// this is guideline or format, how the command should be invoked
	Use: "Compress [inputFile] [outputFile]",
	// short provides brief description of command's use, this description will be shown
	// in the help output when the command is listed among other command
	Short: "Compress video to reduce the files size",
	// defines how many arguments are required for command to execute
	Args: cobra.ExactArgs(2),
	// run contains the func that will be executed when command is called
	// args parameter is a slice of strings (args[0] = inputfile, args[1] = outputfile)
	Run: func(cmd *cobra.Command, args []string) {
		inputFile := args[0]
		outputFile := args[1]
		err := compressVideo(inputFile, outputFile)
		if err != nil {
			fmt.Printf("Failed to compress video: %v\n", err)
		}
	},
}

// init is a special function which is automatically executed when the
// package is loaded, it is here used to register the compressCmd 
// command with rootCmd
func init(){
	rootCmd.AddCommand(compressCmd)
}

func compressVideo(input, output string) error {
	// exec.Command => used to create a new command, which will run the ffmpeg executable
	// ffmpeg -i <input> -vcodec libx264 -crf 28 <output>

	// ffmpeg => Command line tool being executed

	// -i <input> Specifies the input file

	// -vcodec libx264 => Sets the video codec to libx264, which is commonly used for H.264 encoding and is
	// known for achieving a good balance between compression and quality

	// -crf 28 => Constant Rate Factor (CRF) controls the ouput quality
	// lower values result in higher quality (But larger file size)
	// CRF28 is high compression setting reducing file size

	// <output> specifies the location to save the compressed output file
	cmd := exec.Command("ffmpeg", "-i", input, "-vcodec", "libx264", "-crf", "28", output)
	return cmd.Run()
}

