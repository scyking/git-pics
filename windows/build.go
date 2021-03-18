package windows

import (
	"log"
)

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func Build() (*walk.MainWindow, error) {
	var mainWindow *walk.MainWindow
	var treeView *walk.TreeView
	var hSplitter *walk.Splitter
	var vSplitter *walk.Splitter
	var scroll *walk.ScrollView
	var te *walk.TextEdit

	treeModel, err := NewDirectoryTreeModel()
	if err != nil {
		log.Fatal(err)
	}

	if err := (MainWindow{
		AssignTo: &mainWindow,
		Title:    "GPics",
		MinSize:  Size{600, 400},
		Layout:   HBox{MarginsZero: true},
		Children: []Widget{
			HSplitter{
				AssignTo: &hSplitter,
				Children: []Widget{
					TreeView{
						AssignTo: &treeView,
						Model:    treeModel,
						OnCurrentItemChanged: func() {
							dir := treeView.CurrentItem().(*Directory)
							log.Println("path now :", dir.Path())
							if err := te.SetText(dir.Path()); err != nil {
								log.Fatal(err)
							}
							ClearWidgets(scroll)
							AddImageViewWidgets(dir.Path(), scroll)
						},
					},
					VSplitter{
						StretchFactor: 5,
						AssignTo:      &vSplitter,
						Children: []Widget{
							ScrollView{
								AssignTo:      &scroll,
								Name:          "Pictures",
								StretchFactor: 5,
								Layout:        Grid{Columns: 2},
								Children:      []Widget{},
							},
							TextEdit{
								AssignTo: &te,
								ReadOnly: true,
								Text:     "test",
							},
							PushButton{
								Text: "Copy",
								OnClicked: func() {
									if err := walk.Clipboard().SetText(te.Text()); err != nil {
										log.Fatal(err)
									}
								},
							},
						},
					},
				},
			},
		},
	}.Create()); err != nil {
		return nil, err
	}

	if err := vSplitter.SetFixed(scroll, true); err != nil {
		return nil, err
	}
	return mainWindow, nil
}
