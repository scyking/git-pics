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
							log.Println("before children:", scroll.Children())
							ImageViewWidgets(dir.Path(), scroll)
							log.Println("after children:", scroll.Children())
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

func ImageViewWidgets(path string, parent walk.Container) {

	if err := walk.Resources.SetRootDirPath(path); err != nil {
		log.Fatal(err)
	}

	ClearWidgets(parent)

	names := ImageFileNames(path)
	builder := NewBuilder(parent)

	for _, name := range names {
		iv := ImageView{
			Image:   name,
			Margin:  10,
			MinSize: Size{120, 120},
			MaxSize: Size{120, 120},
			Mode:    ImageViewModeZoom,
		}

		if err := iv.Create(builder); err != nil {
			log.Fatal(err)
		}
	}

}

func ClearWidgets(parent walk.Container) {
	widgets := parent.Children()
	if widgets != nil {
		parent.SetSuspended(true)
		defer parent.SetSuspended(false)

		for i := widgets.Len() - 1; i >= 0; i-- {
			widgets.At(i).Dispose()
		}

		if err := widgets.Clear(); err != nil {
			log.Fatal(err)
		}
	}
	log.Println("clear ok!")
}
