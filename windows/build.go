package windows

import (
	"log"
	"path/filepath"
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
							ImageViewWidgets(dir.Path(), scroll)
						},
					},
					VSplitter{
						StretchFactor: 5,
						AssignTo:      &vSplitter,
						Children: []Widget{
							ScrollView{
								AssignTo: &scroll,
								Name:     "Pictures",
								Layout:   Grid{Columns: 2},
								Children: []Widget{
									ImageView{
										MaxSize: Size{120, 120},
										Image:   "1615799620(1).png",
										Margin:  10,
										Mode:    ImageViewModeZoom,
									},
									ImageView{
										MaxSize: Size{120, 120},
										Image:   "1615799620(1).png",
										Margin:  10,
										Mode:    ImageViewModeZoom,
									},
									ImageView{
										MaxSize: Size{120, 120},
										Image:   "1615799620(1).png",
										Margin:  10,
										Mode:    ImageViewModeZoom,
									},
								},
							},
							TextEdit{
								AssignTo: &te,
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

	return mainWindow, nil
}

func ImageViewWidgets(path string, parent walk.Container) {

	names := ImageFileNames(path)

	if parent.Children() != nil {
		if err := parent.Children().Clear(); err != nil {
			log.Fatal(err)
		}
	}

	for i, name := range names {
		log.Println(i, ". pic：", name)

		img, err := walk.NewImageView(parent)
		if err != nil {
			log.Fatal(err)
		}

		if err := img.SetMargin(10); err != nil {
			log.Fatal(err)
		}

		filePath := filepath.Join(path, name)

		log.Println("image path:", filePath)

		image, err := walk.NewImageFromFileForDPI(filePath, 96)
		if err != nil {
			log.Fatal(err)
		}
		if err := img.SetImage(image); err != nil {
			log.Fatal(err)
		}

		img.SetName(name)
		img.SetMode(walk.ImageViewModeZoom)

		size := walk.Size{120, 120}
		if err := img.SetSize(size); err != nil {
			log.Fatal(err)
		}
	}
}
