package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/urizennnn/ScriptPad/Layout"
	"github.com/urizennnn/ScriptPad/Menu"
)

func main() {
	a := app.New()
	w := a.NewWindow("Script Pad")
	// w.Resize(fyne.NewSize(1024, 800))
	Menu.Init()

	// a.Settings().SetTheme(theme.DarkTheme())

	w.SetMainMenu(Menu.AppMenu)
	w.SetContent(Layout.MakeGui())
	w.ShowAndRun()
}
