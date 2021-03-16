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
	var treeView *walk.TreeView
	var vSplitter *walk.Splitter

	var imageViewWidgets []Widget

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
				Children: []Widget{
					TreeView{
						AssignTo: &treeView,
						Model:    treeModel,
						OnCurrentItemChanged: func() {
							dir := treeView.CurrentItem().(*Directory)
							if err := walk.Resources.SetRootDirPath(dir.Path()); err != nil {
								log.Fatal(err)
							}
							imageViewWidgets = ImageViewWidgets(dir.Path())
						},
					},
					VSplitter{
						AssignTo: &vSplitter,
						Children: imageViewWidgets,
					},
				},
			},
		},
	}.Create()); err != nil {
		log.Fatal(err)
	}

	return mainWindow
}

func ImageViewWidgets(filePath string) []Widget {
	var widgets []Widget

	names := DirFiles(filePath)

	for _, name := range names {
		widgets = append(widgets,
			ImageView{
				Image:  name,
				Margin: 10,
			},
		)
	}
	return widgets
}
