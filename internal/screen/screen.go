package screen

import (
	"github.com/GkIgor/xlsreader/cmd"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func InitScreen() {
	box := tview.NewBox().SetBorder(true).SetTitle("XLS Reader").SetBorderAttributes(tcell.AttrBlink)

	cmd.Execute()

	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}

func Cell(text string) *tview.TableCell {
	return tview.NewTableCell(text).SetAlign(tview.AlignCenter)
}
