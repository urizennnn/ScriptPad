package Toolbar

import "fyne.io/fyne/v2"

var ExtendChildMenu = fyne.NewMenuItem("Terminal", nil)
var Help = fyne.NewMenuItem("Help", nil)

var TerminalChild = fyne.NewMenu(
    "",
    fyne.NewMenuItem("New Terminal", nil),
    fyne.NewMenuItem("Split Terminal", nil),
    fyne.NewMenuItem("Run Task", nil),
    fyne.NewMenuItem("Run Build Task", nil),
    fyne.NewMenuItem("Run Active Task", nil),
    fyne.NewMenuItem("Run Selected Task", nil),
    fyne.NewMenuItem("Show Running Task", nil),
    fyne.NewMenuItem("Restart Running Task", nil),
    fyne.NewMenuItem("Terminate Task", nil),
    fyne.NewMenuItem("Configure Task", nil),
    fyne.NewMenuItem("Configure Default Build Task", nil),
)

var helpChild = fyne.NewMenu(
    "",
    fyne.NewMenuItem("Welcome", nil),
    fyne.NewMenuItem("Show All Commands", nil),
    fyne.NewMenuItem("Documentation", nil),
    fyne.NewMenuItem("Editor Playground", nil),
    fyne.NewMenuItem("Show Release Notes", nil),
    fyne.NewMenuItem("Keyboard Shortcuts", nil),
    fyne.NewMenuItem("Video Tutorials", nil),
    fyne.NewMenuItem("Tips and Tricks", nil),
    fyne.NewMenuItem("Join Us On Youtube", nil),
    fyne.NewMenuItem("Search Feature Request", nil),
    fyne.NewMenuItem("Report Issue", nil),
    fyne.NewMenuItem("View License", nil),
    fyne.NewMenuItem("Privacy Statement", nil),
    fyne.NewMenuItem("Toggle Developer Tools", nil),
    fyne.NewMenuItem("Check for Updates", nil),
    fyne.NewMenuItem("About", nil),
)

func ExtendChild() {
    ExtendChildMenu.ChildMenu = TerminalChild
    Help.ChildMenu = helpChild
}
