package Toolbar

import "fyne.io/fyne/v2"

var (
    StartDebugging             = fyne.NewMenuItem("Start Debugging", nil)
    RunWithoutDebugging        = fyne.NewMenuItem("Run Without Debugging", nil)
    StopDebugging              = fyne.NewMenuItem("Stop Debugging", nil)
    RestartDebugging           = fyne.NewMenuItem("Restart Debugging", nil)
    OpenConfigurations         = fyne.NewMenuItem("Open Configurations", nil)
    AddConfigurations          = fyne.NewMenuItem("Add Configurations", nil)
    StepOver                   = fyne.NewMenuItem("Step Over", nil)
    StepInto                   = fyne.NewMenuItem("Step Into", nil)
    StepOut                    = fyne.NewMenuItem("Step Out", nil)
    Continue                   = fyne.NewMenuItem("Continue", nil)
    ToggleBreakpoint           = fyne.NewMenuItem("Toggle Breakpoint", nil)
    NewBreakpoint              = fyne.NewMenuItem("New Breakpoint", nil)
    EnableAllBreakpoint        = fyne.NewMenuItem("Enable All Breakpoint", nil)
    DisableAllBreakpoint       = fyne.NewMenuItem("Disable All Breakpoint", nil)
    RemoveAllBreakpoint        = fyne.NewMenuItem("Remove All Breakpoint", nil)
    InstallAdditionalDebuggers = fyne.NewMenuItem("Install Additional Debuggers", nil)
)
