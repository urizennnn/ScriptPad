package Toolbar

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
)

func CreateToolbar() fyne.CanvasObject {
	file := widget.NewToolbar(
		widget.NewToolbarAction(
			theme.ContentCopyIcon(), func() {

			},
		),
	)
	search := widget.NewToolbar(
		widget.NewToolbarAction(
			theme.SearchIcon(), func() {

			},
		),
	)

	git := LoadGit()
	gitToolbar := widget.NewToolbar(git)
	debug := LoadDebug()
	debugToolbar := widget.NewToolbar(debug)

	extensions := widget.NewToolbar(widget.NewToolbarAction(theme.GridIcon(), func() {}))
	return container.NewVBox(file, search, gitToolbar, extensions, debugToolbar)
}

func LoadGit() *widget.ToolbarAction {
	imageResource, err := fyne.LoadResourceFromPath("images/Git.png")
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	toolbarAction := widget.NewToolbarAction(
		imageResource, func() {

		},
	)

	return toolbarAction
}
func LoadDebug() *widget.ToolbarAction {
	imageResource, err := fyne.LoadResourceFromPath("images/debug.png")
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	toolbarAction := widget.NewToolbarAction(
		imageResource, func() {

		},
	)

	return toolbarAction
}

func Utility() fyne.CanvasObject {
	settings := widget.NewToolbar(widget.NewToolbarAction(theme.SettingsIcon(), func() {}))
	profile := widget.NewToolbar(widget.NewToolbarAction(theme.AccountIcon(), func() {}))

	return container.NewVBox(profile, settings)

}
