package cmd

import (
	"fmt"
	"os"

	"github.com/GkIgor/xlsreader/internal/reader"
	"github.com/spf13/cobra"
)

var filepath string
var sheetName string

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read an XLSX file and display its content",
	Run: func(cmd *cobra.Command, args []string) {
		if filepath == "" {
			fmt.Println("You must specify a file using --file")
			os.Exit(1)
			return
		}
		reader.Ready(filepath, sheetName)
	},
}

func init() {
	readCmd.Flags().StringVarP(&filepath, "file", "f", "", "Path to the XLSX file")
	readCmd.Flags().StringVarP(&filepath, "sheet", "s", "", "Sheet name (optional)")
	rootCmd.AddCommand(readCmd)
}
