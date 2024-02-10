package Toolbar

import "fyne.io/fyne/v2"

var (
	SelectAll           = fyne.NewMenuItem("Select All", nil)
	ExpandSelection     = fyne.NewMenuItem("Expand Selection", nil)
	ShrinkSelection     = fyne.NewMenuItem("Shrink Selection", nil)
	CopyLineUp          = fyne.NewMenuItem("Copy Line Up", nil)
	CopyLineDown        = fyne.NewMenuItem("Copy Line Down", nil)
	MoveLineUp          = fyne.NewMenuItem("Move Line Up", nil)
	MoveLineDown        = fyne.NewMenuItem("Move Line Down", nil)
	DuplicateSelection  = fyne.NewMenuItem("Duplicate Selection", nil)
	AddCursorAbove      = fyne.NewMenuItem("Add Cursor Above", nil)
	AddCursorBelow      = fyne.NewMenuItem("Add Cursor Below", nil)
	AddCursorToLindEnds = fyne.NewMenuItem("Add Cursor To Line Ends", nil)
	AddNextOccurrence   = fyne.NewMenuItem("Add Previous Occurrence", nil)
	SelectAllOccurrence = fyne.NewMenuItem("Select All Occurrence", nil)
	SwitchToAlt         = fyne.NewMenuItem("Switch To Alt For Multi-Cursor Selection", nil)
	ColumnSelectionMode = fyne.NewMenuItem("Column Selection Mode", nil)
)
