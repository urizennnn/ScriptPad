package Toolbar

import "fyne.io/fyne/v2"

var (
	Back                  = fyne.NewMenuItem("Back", nil)
	Forward               = fyne.NewMenuItem("Forward", nil)
	LastEditLocation      = fyne.NewMenuItem("Last Edit Location", nil)
	SwitchEditor          = fyne.NewMenuItem("Switch Editor", nil)
	SwitchGroup           = fyne.NewMenuItem("Switch Group", nil)
	GoToFile              = fyne.NewMenuItem("Go To File", nil)
	GoToSymbolInWorkSpace = fyne.NewMenuItem("Go To Symbol In WorkSpace", nil)
	GoToSymbolInEditor    = fyne.NewMenuItem("Go To Symbol In Editor", nil)
	GoToDefinition        = fyne.NewMenuItem("Go To Definition", nil)
	GoToDeclaration       = fyne.NewMenuItem("Go To Declaration", nil)
	GoToTypeDefinition    = fyne.NewMenuItem("Go To Type Definition", nil)
	GoToImplementation    = fyne.NewMenuItem("Go To Implementation", nil)
	GoToReferences        = fyne.NewMenuItem("Go To References", nil)
	GoToLineColumn        = fyne.NewMenuItem("Go To Line/Column", nil)
	GoToBrackets          = fyne.NewMenuItem("Go To Brackets", nil)
	NextProblem           = fyne.NewMenuItem("Next Problem", nil)
	PreviousProblem       = fyne.NewMenuItem("Previous Problem", nil)
	NextChange            = fyne.NewMenuItem("Next Change", nil)
	PreviousChange        = fyne.NewMenuItem("Previous Change", nil)
)

var editorChild = fyne.NewMenu(
	"",
	fyne.NewMenuItem("New Editor", nil),
	fyne.NewMenuItem("Previous Editor", nil),
	fyne.NewMenuItem("Next Used Editor", nil),
	fyne.NewMenuItem("Previous Used Editor", nil),
	fyne.NewMenuItem("Next Editor In Group", nil),
	fyne.NewMenuItem("Previous Editor In Group", nil),
	fyne.NewMenuItem("Next Used Editor In Group", nil),
	fyne.NewMenuItem("Previous Used Editor In Group", nil),
)
var groupChild = fyne.NewMenu(
	"",
	fyne.NewMenuItem("Group 1", nil),
	fyne.NewMenuItem("Group 2", nil),
	fyne.NewMenuItem("Group 3", nil),
	fyne.NewMenuItem("Group 4", nil),
	fyne.NewMenuItem("Group 5", nil),
	fyne.NewMenuItem("Next Group", nil),
	fyne.NewMenuItem("Previous Group", nil),
	fyne.NewMenuItem("Group Left", nil),
	fyne.NewMenuItem("Group Right", nil),
	fyne.NewMenuItem("Group Up", nil),
	fyne.NewMenuItem("Group Below", nil),
)

func GoChild() {
	SwitchEditor.ChildMenu = editorChild
	SwitchGroup.ChildMenu = groupChild
}
