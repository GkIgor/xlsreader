package xlsx

import (
	"fmt"
	"os"

	"github.com/GkIgor/xlsreader/internal/models"
)

func OpenXLSX(coreSheet *models.CoreSheet) {
	fmt.Printf("Opening %s\n", coreSheet.FilePath)
	rows, err := coreSheet.File.GetRows(coreSheet.DefaultSheetname)

	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	if len(rows) == 0 {
		fmt.Println("No rows found")
	}

	for _, row := range rows[0:10] {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}
