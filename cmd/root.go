package cmd

import (
	"fmt"
	"os"

	"github.com/GkIgor/xlsreader/global"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   global.APP_NAME,
	Short: "A CLI tool to read and process XLS/XLSX files",
	Long:  fmt.Sprintf("%s%s%s\n %s - %s\n  `xlsreader` is a lightweight and efficient command-line tool for reading and processing `.xls` and `.xlsx` files on Linux systems. It allows users to quickly view spreadsheet data, extract specific sheets, and export content in various formats.", global.Green, global.APP_NAME, global.Reset, global.APP_VERSION, global.APP_FLAG),
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
