package cmd

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"sync"

	"github.com/spf13/cobra"
)

var batchCompresscmd = &cobra.Command{
	Use: "batch-compress [directory]",
	Short: "Compress all videos in a directory",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dir := args[0]
		var wg sync.WaitGroup

		filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
			if !info.IsDir() && filepath.Ext(info.Name()) == ".mp4" {
				wg.Add(1)
				go func(filePath string){
					defer wg.Done()
					output := fmt.Sprintf("%s_compressed.mp4", filePath)
					err := compressVideo(filePath, output)
					if err != nil {
						fmt.Printf("Failed to compress %s: %v\n", filePath, err)
					}
				}(path)
			}
			return nil
		})
		wg.Wait()
		fmt.Println("Batch compression completed!")
	},
}

func init(){
	rootCmd.AddCommand(batchCompresscmd)
}