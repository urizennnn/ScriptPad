package Layout

import (
    `fmt`
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    `fyne.io/fyne/v2/data/binding`
    `fyne.io/fyne/v2/dialog`
    "fyne.io/fyne/v2/theme"
    "fyne.io/fyne/v2/widget"
    `github.com/urizennnn/ScriptPad/Toolbar`
    "image/color"
    `log`
)

type scriptLayout struct {
    content, left, divider fyne.CanvasObject
}

type Gui struct {
    Win       fyne.Window
    directory *widget.Label
    tree      binding.URITree
}

func newScript(left, content, divider fyne.CanvasObject) fyne.Layout {
    return &scriptLayout{left: left, content: content, divider: divider}
}

func (L *scriptLayout) Layout(_ []fyne.CanvasObject, size fyne.Size) {
    
    dividerWidth := L.left.MinSize().Width
    thick := theme.SeparatorThicknessSize()
    L.divider.Move(fyne.NewPos(dividerWidth, thick))
    L.divider.Resize(fyne.NewSize(2.453, size.Height))
    L.content.Move(fyne.NewPos(dividerWidth, 0))
    L.content.Resize(fyne.NewSize(size.Width, size.Height))
    
}

func (L *scriptLayout) MinSize([]fyne.CanvasObject) fyne.Size {
    return fyne.NewSize(900, 400)
}

func (g *Gui) MakeGui() fyne.CanvasObject {
    g.directory = widget.NewLabel("Open a file")
    g.tree = binding.NewURITree()
    files := widget.NewTreeWithData(
        g.tree, func(branch bool) fyne.CanvasObject {
            return widget.NewLabel("Testing files")
        }, func(data binding.DataItem, branch bool, obj fyne.CanvasObject) {
            l := obj.(*widget.Label)
            get, err := data.(binding.URI).Get()
            if err != nil {
                return
            }
            name := get.Name()
            l.SetText(name)
        },
    )
    
    left := container.NewStack(CreateToolbar(g.Win, &Gui{Win: g.Win}), files)
    
    //tree := g.tree
    content := container.NewStack(
        canvas.NewRectangle(
            color.Gray{Y: 0xee},
        ), container.NewCenter(g.directory),
    )
    
    divider := widget.NewSeparator()
    objs := []fyne.CanvasObject{content, left, divider}
    return container.New(newScript(left, content, divider), objs...)
}
func (g *Gui) OpenFolder(w fyne.Window) {
    fmt.Println("open")
    dialog.ShowFolderOpen(
        func(dir fyne.ListableURI, err error) {
            if err != nil {
                log.Printf("Error opening folder: %v\n", err)
                dialog.ShowError(err, w)
                return
            }
            
            if dir == nil {
                log.Println("Folder selection canceled")
                return
            }
            g.Open(dir)
        }, w,
    )
}

func (g *Gui) Open(dir fyne.ListableURI) {
    name := dir.Name()
    g.directory.SetText(name)
    g.Win.SetTitle(name)
    
    list, err := dir.List()
    if err != nil {
        dialog.ShowError(err, g.Win)
        return
    }
    for _, file := range list {
        errors := g.tree.Append(binding.DataTreeRootID, file.String(), file)
        if errors != nil {
            dialog.ShowError(errors, g.Win)
            return
        }
    }
}

func CreateToolbar(w fyne.Window, gui *Gui) fyne.CanvasObject {
    file := widget.NewToolbar(
        widget.NewToolbarAction(
            theme.ContentCopyIcon(), func() {
                gui.OpenFolder(w)
            },
        ),
    )
    search := widget.NewToolbar(
        widget.NewToolbarAction(
            theme.SearchIcon(), func() {
                // Add functionality for search if needed
            },
        ),
    )
    
    git := Toolbar.LoadGit()
    gitToolbar := widget.NewToolbar(git)
    debug := Toolbar.LoadDebug()
    debugToolbar := widget.NewToolbar(debug)
    
    extensions := widget.NewToolbar(
        widget.NewToolbarAction(
            theme.GridIcon(), func() {
                // Add functionality for extensions if needed
            },
        ),
    )
    
    return container.NewVBox(file, search, gitToolbar, extensions, debugToolbar, Toolbar.Utility())
}
