package windows

import (
	"log"
)

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func Build() *walk.MainWindow {
	var mainWindow *walk.MainWindow
	var splitter *walk.Splitter
	var treeView *walk.TreeView

	treeModel, err := NewDirectoryTreeModel()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("开始构建")
	if err := (MainWindow{
		AssignTo: &mainWindow,
		Title:    "GPics",
		MinSize:  Size{600, 400},
		Layout:   VBox{},
		Children: []Widget{
			HSplitter{
				AssignTo: &splitter,
				Children: []Widget{
					TreeView{
						AssignTo: &treeView,
						Model:    treeModel,
					},
				},
			},
		},
	}.Create()); err != nil {
		log.Fatal(err)
	}
	log.Println("构建结束")
	return mainWindow
}
