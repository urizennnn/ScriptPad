package main

import (
    "fmt"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
    "github.com/urizennnn/ScriptPad/Explorer"
)

func main() {
    a := app.New()
    w := a.NewWindow("Script Pad")
    w.Resize(fyne.NewSize(1200, 800))
    icon, err := fyne.LoadResourceFromPath("C:/Users/Victor/Desktop/Urizen's Text  Editor/images/chinmay-b-fd9mIBluHkA-unsplash.jpg")
    if err != nil {
        fmt.Print(err)
        a.Quit()
        return
    }
    
    a.SetIcon(icon)
    
    ExplorerSection := widget.NewList(
        func() int {
            return len(Explorer.FieldItem)
        }, func() fyne.CanvasObject {
            return widget.NewLabel("")
        }, func(id widget.ListItemID, object fyne.CanvasObject) {
            object.(*widget.Label).SetText(Explorer.FieldItem[id])
        },
    )
    
    // Set a max size for the split
    maxSize := fyne.NewSize(8, 6)
    split := container.NewHSplit(ExplorerSection, container.NewMax())
    split.Resize(maxSize)
    split.Offset = 0.
    
    w.SetContent(split)
    w.ShowAndRun()
}
