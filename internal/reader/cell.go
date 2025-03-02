package reader

import (
	"fmt"

	exc "github.com/xuri/excelize/v2"
)

func PrintCells(file *exc.File, sheetname string) {
	rows, err := file.GetRows(sheetname)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}

}
