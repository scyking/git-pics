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
	var splitter *walk.Splitter
	var scroll *walk.ScrollView

	//var imageViewWidgets []Widget

	walk.Resources.SetRootDirPath("C:\\workspace\\test")

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
				AssignTo: &splitter,
				Children: []Widget{
					TreeView{
						AssignTo: &treeView,
						Model:    treeModel,
						OnCurrentItemChanged: func() {
							dir := treeView.CurrentItem().(*Directory)
							log.Println("path now :", dir.Path())
							//if err := walk.Resources.SetRootDirPath(dir.Path()); err != nil {
							//	log.Fatal(err)
							//}
							//imageViewWidgets = ImageViewWidgets(dir.Path())

						},
					},
					ScrollView{
						AssignTo:      &scroll,
						Name:          "Pictures",
						StretchFactor: 2,
						Layout:        Grid{Columns: 1},
						Children: []Widget{
							ImageView{
								MaxSize: Size{120, 120},
								Image:   "1615799620(1).png",
								Margin:  10,
								Mode:    ImageViewModeZoom,
							},
						},
					},
				},
			},
		},
	}.Create()); err != nil {
		log.Fatal(err)
	}

	if err := splitter.SetFixed(scroll, true); err != nil {
		log.Fatal(err)
	}

	return mainWindow
}

func ImageViewWidgets(filePath string) []Widget {
	var widgets []Widget

	names := DirFiles(filePath)

	for i, name := range names {
		log.Println(i, ". pic：", name)
		widgets = append(widgets,
			ImageView{
				MaxSize: Size{120, 120},
				Image:   name,
				Margin:  10,
				Mode:    ImageViewModeZoom,
			},
		)
	}
	return widgets
}
