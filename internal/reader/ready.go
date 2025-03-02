package reader

import (
	"fmt"
	"os"
)

func Ready(filepath string, sheetName string) {
	fmt.Println("Ready?")

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	Read(filepath, sheetName)
}
