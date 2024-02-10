package Toolbar

import "fyne.io/fyne/v2"

var (
	NewTextFile           = fyne.NewMenuItem("New Text File", nil)
	NewFile               = fyne.NewMenuItem("New File", nil)
	NewWindow             = fyne.NewMenuItem("New Window", nil)
	OpenFile              = fyne.NewMenuItem("Open File", nil)
	OpenFolder            = fyne.NewMenuItem("Open Folder", nil)
	OpenWorkSpaceFromFile = fyne.NewMenuItem("Open Work Space From File", nil)
	OpenRecent            = fyne.NewMenuItem("Open Recent", nil)
	AddFolderToWorkSpace  = fyne.NewMenuItem("Add  Folder To Workspace", nil)
	SaveWorkSpaceAs       = fyne.NewMenuItem("Save WorkSpace as", nil)
	DuplicateWorkSpace    = fyne.NewMenuItem("Duplicate Workspace", nil)
	Save                  = fyne.NewMenuItem("Save", nil)
	SaveAs                = fyne.NewMenuItem("Save As", nil)
	SaveAll               = fyne.NewMenuItem("Save All", nil)
	Share                 = fyne.NewMenuItem("Share", nil)
	AutoSave              = fyne.NewMenuItem("Auto Save", nil)
	Preferences           = fyne.NewMenuItem("Preferences", nil)
	RevertFile            = fyne.NewMenuItem("Revert File", nil)
	CloseEditor           = fyne.NewMenuItem("Close Editor", nil)
	CloseWindow           = fyne.NewMenuItem("Close Window", nil)
	Exit                  = fyne.NewMenuItem("Exit", nil)
)

// OpenRecentChildMenu Child menu for OpenRecent
var OpenRecentChildMenu = fyne.NewMenu(
	"",
	fyne.NewMenuItem("More", nil),
)
var ShareMenu = fyne.NewMenu(
	"",
	fyne.NewMenuItem("Export File", nil),
	fyne.NewMenuItem("Import File", nil),
)

var PreferencesMenu = fyne.NewMenu(
	"",
	fyne.NewMenuItem("Profile", nil),
	fyne.NewMenuItem("Settings", nil),
	fyne.NewMenuItem("Extensions", nil),
	fyne.NewMenuItem("Keyboard Shortcuts", nil),
	fyne.NewMenuItem("Configure User Snippets", nil),
	fyne.NewMenuItem("User Tasks", nil),
	fyne.NewMenuItem("Theme", nil),
	fyne.NewMenuItem("Online Services", nil),
)

func Init() {
	OpenRecent.ChildMenu = OpenRecentChildMenu
	Share.ChildMenu = ShareMenu
	Preferences.ChildMenu = PreferencesMenu
}
