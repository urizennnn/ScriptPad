package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    `fyne.io/x/fyne/theme`
    "github.com/urizennnn/ScriptPad/Menu"
    `github.com/urizennnn/ScriptPad/internal/Layout`
)

func main() {
    var a = app.New()
    a.Settings().SetTheme(theme.AdwaitaTheme())
    var W = a.NewWindow("Script Pad")
    var Ui = &Layout.Gui{Win: W}
    
    W.Resize(fyne.NewSize(900, 700))
    Menu.Init()
    
    // a.Settings().SetTheme(theme.DarkTheme())
    W.SetMainMenu(Menu.AppMenu)
    W.SetContent(Ui.MakeGui())
    
    W.ShowAndRun()
}
