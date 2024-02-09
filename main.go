package main

import (
	"fmt"
	"image/color"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Script Pad")
	w.Resize(fyne.NewSize(400, 400))
	icon,err:=fyne.LoadResourceFromPath("./images/chinmay-b-fd9mIBluHkA-unsplash.jpg")
	if err != nil {
		panic(err)
	}
	
	a.SetIcon(icon)

	setTheme := widget.NewButton("Dark Theme", func() {
		a.Settings().SetTheme(theme.DarkTheme())
	})

	fileExplorer := widget.NewButton("File Explorer", func() {
		fmt.Println("File Explorer Section")
	})
	Textarea := canvas.NewText("Text Area", color.Black)
	toolbar := widget.NewButton("Toolbar", func() {
		fmt.Print("Toolbar")
	})

	w.SetContent(
		container.NewVBox(toolbar, setTheme,
			container.NewHSplit(
				fileExplorer, Textarea,
			),
		),
	)
	w.ShowAndRun()
}
