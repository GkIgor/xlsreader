package xlsx

import (
	"fmt"

	"github.com/GkIgor/xlsreader/internal/models"
)

func OpenXLSX(coreSheet *models.CoreSheet) {
	fmt.Printf("Opening %s\n", coreSheet.FilePath)
}
