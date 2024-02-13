package Layout

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/urizennnn/ScriptPad/Toolbar"
	"image/color"
	"log"
)

type scriptLayout struct {
	content      fyne.CanvasObject
	left, bottom *widget.Toolbar
}

func newScript(content fyne.CanvasObject, bottom, left *widget.Toolbar) fyne.Layout {
	return &scriptLayout{bottom: bottom, left: left, content: content}
}

func (l *scriptLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	//topHeight := l.top.MinSize().Height
	//sideWidth := l
	log.Println("Size", size)
}
func (l *scriptLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.Size{Width: 1024, Height: 500}
}

func MakeGui() fyne.CanvasObject {

	content := canvas.NewRectangle(
		color.Gray{
			Y: 0xee,
		},
	)

	return container.NewBorder(nil, Toolbar.Utility(), Toolbar.CreateToolbar(), nil, content)
}
