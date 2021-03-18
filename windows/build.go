package windows

import (
	"log"
)

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func Build() (*walk.MainWindow, error) {
	var mw *walk.MainWindow
	var tv *walk.TreeView
	var hs *walk.Splitter
	var vs *walk.Splitter
	var sv *walk.ScrollView
	var te *walk.TextEdit

	treeModel, err := NewDirectoryTreeModel()
	if err != nil {
		log.Fatal(err)
	}

	if err := (MainWindow{
		AssignTo: &mw,
		Title:    "GPics",
		MinSize:  Size{600, 400},
		Layout:   HBox{MarginsZero: true},
		Children: []Widget{
			HSplitter{
				AssignTo: &hs,
				Children: []Widget{
					TreeView{
						AssignTo: &tv,
						Model:    treeModel,
						OnCurrentItemChanged: func() {
							dir := tv.CurrentItem().(*Directory)
							log.Println("path now :", dir.Path())
							if err := te.SetText(dir.Path()); err != nil {
								log.Fatal(err)
							}
							ClearWidgets(sv)
							AddImageViewWidgets(dir.Path(), sv)
						},
					},
					VSplitter{
						StretchFactor: 5,
						AssignTo:      &vs,
						Children: []Widget{
							ScrollView{
								AssignTo:      &sv,
								Name:          "Pictures",
								StretchFactor: 5,
								Layout: Flow{
									MarginsZero: true,
									Spacing:     5,
									Alignment:   AlignHNearVCenter,
								},
								Children: []Widget{},
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

	if err := vs.SetFixed(sv, true); err != nil {
		return nil, err
	}
	return mw, nil
}
