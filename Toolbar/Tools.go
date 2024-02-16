package Toolbar

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/theme"
    main `fyne.io/fyne/v2/widget`
    "log"
)

func CreateToolbar() fyne.CanvasObject {
    log.Println("created")
    
    file := main.NewToolbar(
        main.NewToolbarAction(
            theme.ContentCopyIcon(), func() {
                log.Println("testing")
            },
        ),
    )
    search := main.NewToolbar(
        main.NewToolbarAction(
            theme.SearchIcon(), func() {
            
            },
        ),
    )
    
    git := LoadGit()
    gitToolbar := main.NewToolbar(git)
    debug := LoadDebug()
    debugToolbar := main.NewToolbar(debug)
    
    extensions := main.NewToolbar(main.NewToolbarAction(theme.GridIcon(), func() {}))
    return container.NewVBox(file, search, gitToolbar, extensions, debugToolbar, Utility())
}

func LoadGit() *main.ToolbarAction {
    imageResource, err := fyne.LoadResourceFromPath("images/Git.png")
    if err != nil {
        log.Fatalln(err)
        return nil
    }
    
    toolbarAction := main.NewToolbarAction(
        imageResource, func() {
        
        },
    )
    
    return toolbarAction
}
func LoadDebug() *main.ToolbarAction {
    imageResource, err := fyne.LoadResourceFromPath("images/debug.png")
    if err != nil {
        log.Fatalln(err)
        return nil
    }
    
    toolbarAction := main.NewToolbarAction(
        imageResource, func() {
        
        },
    )
    
    return toolbarAction
}

func Utility() fyne.CanvasObject {
    settings := main.NewToolbar(main.NewToolbarAction(theme.SettingsIcon(), func() {}))
    profile := main.NewToolbar(main.NewToolbarAction(theme.AccountIcon(), func() {}))
    
    return container.NewVBox(profile, settings)
    
}
