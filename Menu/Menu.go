package Menu

import "fyne.io/fyne/v2"

var File = fyne.NewMenu(
	"File", NewTextFile, NewFile, NewWindow, OpenFile, OpenFolder, OpenWorkSpaceFromFile, OpenRecent,
	AddFolderToWorkSpace, SaveWorkSpaceAs, DuplicateWorkSpace, Save, SaveAs, SaveAll, Share, AutoSave, Preferences,
	RevertFile, CloseEditor, CloseWindow, Exit,
)

var Edit = fyne.NewMenu(
	"Edit", Undo, Redo, Cut, Copy, Paste, Find, Replace, FindFiles, ReplaceFiles, ToggleLineComment, ToggleBlockComment,
)
var View = fyne.NewMenu(
	"View", CommandPalette, OpenView, Appearance, EditorLayout,
	Explorer, Search, SourceControl, Run, Extensions, Problems, Output,
	DebugConsole, Terminals, WordWrap,
)
var Selection = fyne.NewMenu(
	"Selection", SelectAll, ExpandSelection, ShrinkSelection,
	CopyLineUp, CopyLineDown, MoveLineUp, MoveLineDown, DuplicateSelection,
	AddCursorAbove, AddCursorBelow, AddCursorToLindEnds, AddNextOccurrence,
	SelectAllOccurrence, SwitchToAlt, ColumnSelectionMode,
)
var Go = fyne.NewMenu(
	"Go", Back, Forward, LastEditLocation, SwitchEditor, SwitchGroup,
	GoToFile, GoToSymbolInWorkSpace, GoToSymbolInEditor, GoToDefinition, GoToDeclaration,
	GoToTypeDefinition, GoToImplementation, GoToReferences, GoToLineColumn, GoToBrackets,
	NextProblem, PreviousProblem, NextChange, PreviousChange,
)

var run = fyne.NewMenu(
	"Run", StartDebugging, RunWithoutDebugging, StopDebugging, RestartDebugging,
	OpenConfigurations, AddConfigurations, StepOver, StepInto, StepOut, Continue,
	ToggleBreakpoint, NewBreakpoint, EnableAllBreakpoint, DisableAllBreakpoint, RemoveAllBreakpoint,
	InstallAdditionalDebuggers,
)
var Extend = fyne.NewMenu("...", ExtendChildMenu, Help)

var AppMenu = fyne.NewMainMenu(File, Edit, Selection, View, Go, run, Extend)

func Init() {
	FileView()
	ViewChild()
	GoChild()
	RunChild()
	ExtendChild()
}
