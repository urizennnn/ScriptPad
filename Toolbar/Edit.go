package Toolbar

import "fyne.io/fyne/v2"

var (
	Undo               = fyne.NewMenuItem("Undo", nil)
	Redo               = fyne.NewMenuItem("Redo", nil)
	Cut                = fyne.NewMenuItem("Cut", nil)
	Copy               = fyne.NewMenuItem("Copy", nil)
	Paste              = fyne.NewMenuItem("Paste", nil)
	Find               = fyne.NewMenuItem("Find", nil)
	Replace            = fyne.NewMenuItem("Replace", nil)
	FindFiles          = fyne.NewMenuItem("Find in Files", nil)
	ReplaceFiles       = fyne.NewMenuItem("Replace Files", nil)
	ToggleLineComment  = fyne.NewMenuItem("Toggle Line Comment", nil)
	ToggleBlockComment = fyne.NewMenuItem("Toggle Block Comment", nil)
)
