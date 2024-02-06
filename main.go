package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")
	// label:=widget.NewLabel("Hello World!")

	btn := widget.NewButton("Testing button", func() {
		fmt.Println("Button working")
	})
	w.SetContent(btn)
	w.ShowAndRun()
}
