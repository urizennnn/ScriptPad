package Menu

import "fyne.io/fyne/v2"

var (
	CommandPalette = fyne.NewMenuItem("Command Palette", nil)
	OpenView       = fyne.NewMenuItem("Open View", nil)
	Appearance     = fyne.NewMenuItem("Appearance", nil)
	EditorLayout   = fyne.NewMenuItem("Editor Layout ", nil)
	Explorer       = fyne.NewMenuItem("Explorer ", nil)
	Search         = fyne.NewMenuItem("Search", nil)
	SourceControl  = fyne.NewMenuItem("Source Control", nil)
	Run            = fyne.NewMenuItem("Run", nil)
	Extensions     = fyne.NewMenuItem("Extensions", nil)
	Problems       = fyne.NewMenuItem("Problems", nil)
	Output         = fyne.NewMenuItem("Output", nil)
	DebugConsole   = fyne.NewMenuItem("Debug Console", nil)
	Terminals      = fyne.NewMenuItem("Terminals", nil)
	WordWrap       = fyne.NewMenuItem("Word Wrap", nil)
)

// Creating ChildMenu for variables

var AppearanceChild = fyne.NewMenu(
	"",
	fyne.NewMenuItem("Full Screen", nil),
	fyne.NewMenuItem("Zen Mode", nil),
	fyne.NewMenuItem("Centered Layout", nil),
	fyne.NewMenuItem("Menu Bar", nil),
	fyne.NewMenuItem("Primary Side Bar", nil),
	fyne.NewMenuItem("Secondary Side Bar", nil),
	fyne.NewMenuItem("Status", nil),
	fyne.NewMenuItem("Panel", nil),
	fyne.NewMenuItem("Move Primary Side Bar Right", nil),
	fyne.NewMenuItem("Activity Bar Position", nil),
	fyne.NewMenuItem("Panel Position", nil),
	fyne.NewMenuItem("Align Plane", nil),
	fyne.NewMenuItem("Tab Bar", nil),
	fyne.NewMenuItem("Editor Actions Position", nil),
	fyne.NewMenuItem("MiniMap", nil),
	fyne.NewMenuItem("Breadcrumbs", nil),
	fyne.NewMenuItem("Sticky Scroll", nil),
	fyne.NewMenuItem("Render Whitespace", nil),
	fyne.NewMenuItem("Render Control Characters", nil),
	fyne.NewMenuItem("Zoom In", nil),
	fyne.NewMenuItem("Zoom Out", nil),
	fyne.NewMenuItem("Reset Control", nil),
)

var EditorChild = fyne.NewMenu(
	"",
	fyne.NewMenuItem("Split Up", nil),
	fyne.NewMenuItem("Split Down", nil),
	fyne.NewMenuItem("Split Left", nil),
	fyne.NewMenuItem("Split Right", nil),
	fyne.NewMenuItem("Split In Group", nil),
	fyne.NewMenuItem("Move Editor Into New Window", nil),
	fyne.NewMenuItem("Copy Editor Into New Window", nil),
	fyne.NewMenuItem("Single", nil),
	fyne.NewMenuItem("Two Columns", nil),
	fyne.NewMenuItem("Three Columns", nil),
	fyne.NewMenuItem("Two Rows", nil),
	fyne.NewMenuItem("Three Rows", nil),
	fyne.NewMenuItem("Grid (2x2)", nil),
	fyne.NewMenuItem("Two Rows Right", nil),
	fyne.NewMenuItem("Two Rows Bottom", nil),
	fyne.NewMenuItem("Flip Layout", nil),
)

func ViewChild() {
	Appearance.ChildMenu = AppearanceChild
	EditorLayout.ChildMenu = EditorChild

}
