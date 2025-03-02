package reader

import (
	"fmt"
	"os"
	fp "path/filepath"

	"github.com/GkIgor/xlsreader/internal/helpers"
	"github.com/GkIgor/xlsreader/internal/models"
	"github.com/GkIgor/xlsreader/internal/reader/xlsx"
	exc "github.com/xuri/excelize/v2"
)

func Read(filepath string, sheetName string) {
	ext := fp.Ext(filepath)

	var coreSheet *models.CoreSheet = new(models.CoreSheet)

	switch ext {
	case ".xls", ".xlsx":

		coreSheet.FilePath = filepath
		coreSheet.DefaultSheetname = sheetName
		OpenExcel(coreSheet)
	}

}

func OpenExcel(coreSheet *models.CoreSheet) {
	f, err := exc.OpenFile(coreSheet.FilePath)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	sheetnames := f.GetSheetList()

	if len(sheetnames) == 0 {
		fmt.Println("No sheets found")
		os.Exit(0)
	}

	if coreSheet.DefaultSheetname != "" {
		_, exists := helpers.MakeSliceSet(sheetnames)[coreSheet.DefaultSheetname]

		if !exists {
			fmt.Printf("Sheet %s not found\n", coreSheet.DefaultSheetname)
			os.Exit(1)
		}
	} else {
		coreSheet.DefaultSheetname = sheetnames[0]
	}

	coreSheet.SheetNames = sheetnames

	xlsx.OpenXLSX(coreSheet)
}
