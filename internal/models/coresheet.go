package models

import exc "github.com/xuri/excelize/v2"

type CoreSheet struct {
	DefaultSheetname string
	SheetNames       []string
	FilePath         string
	File             *exc.File
}
