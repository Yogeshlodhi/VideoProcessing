package utils

import (
	"fmt"
	"os"
)

func FileExists(filename string) bool {
	info, err := os.Stat(filename)

	if os.IsNotExist(err){
		return false
	}
	return !info.IsDir()
}

func CheckErr(err error){
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		os.Exit(1)
	}
}

