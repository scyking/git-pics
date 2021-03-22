package windows

import (
	"gpics/config"
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
	var le *walk.LineEdit

	// 数据绑定
	db := DBSource()

	treeModel, err := NewDirectoryTreeModel()
	if err != nil {
		log.Fatal(err)
	}

	if err := (MainWindow{
		AssignTo: &mw,
		Title:    config.PName,
		MinSize:  Size{600, 400},
		Layout:   HBox{MarginsZero: true},
		DataBinder: DataBinder{
			AutoSubmit: true,
			DataSource: db,
		},
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
							if err := le.SetText(dir.Path()); err != nil {
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
								AssignTo: &sv,
								Name:     "Pictures",
								DataBinder: DataBinder{
									DataSource: db,
								},
								Layout: Flow{
									MarginsZero: true,
									Spacing:     5,
									Alignment:   AlignHNearVCenter,
								},
								Children: []Widget{},
							},
							LineEdit{
								AssignTo: &le,
								ReadOnly: true,
								Text:     "test",
							},
							RadioButtonGroup{
								DataMember: DBTextType,
								Buttons: []RadioButton{
									{
										Name:  "'Markdown' Text",
										Text:  "Markdown",
										Value: Markdown,
										OnClicked: func() {
											//
										},
									},
									{
										Name:  "'HTML' Text",
										Text:  "HTML",
										Value: HTML,
									},
									{
										Name:  "'URL' Text",
										Text:  "URL",
										Value: URL,
									},
									{
										Name:  "'FilePath' Text",
										Text:  "FilePath",
										Value: FilePath,
									},
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
