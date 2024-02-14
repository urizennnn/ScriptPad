package Layout

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/urizennnn/ScriptPad/Toolbar"
	"image/color"
)

//var sideWidth float32 = 220

type scriptLayout struct {
	content, left fyne.CanvasObject
	divider       fyne.CanvasObject
}

func newScript(left, content, divider fyne.CanvasObject) fyne.Layout {
	return &scriptLayout{left: left, content: content, divider: divider}
}

func (L *scriptLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {

	dividerWidth := L.left.MinSize().Width
	thick := theme.SeparatorThicknessSize()
	L.divider.Move(fyne.NewPos(dividerWidth, thick))
	L.divider.Resize(fyne.NewSize(2.453, size.Height))
	L.content.Move(fyne.NewPos(dividerWidth, 0))
	L.content.Resize(fyne.NewSize(size.Width, size.Height))
}

func (L *scriptLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(900, 400)
}

func MakeGui() fyne.CanvasObject {
	left := Toolbar.CreateToolbar()
	//bottom := Toolbar.Utility()
	content := canvas.NewRectangle(
		color.Gray{Y: 0xee},
	)

	divider := widget.NewSeparator()
	objs := []fyne.CanvasObject{content, left, divider}
	return container.New(newScript(left, content, divider), objs...)
}
