package main

import (
	"git-pics/windows"
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

	treeModel, err := windows.NewDirectoryTreeModel()
	if err != nil {
		log.Fatal(err)
	}

	// 数据绑定
	db := make(map[string]string)
	// 设置text type默认类型
	db["tt"] = "b"

	if err := (MainWindow{
		AssignTo: &mw,
		Title:    "GPics",
		MinSize:  Size{600, 400},
		Layout:   HBox{MarginsZero: true},
		DataBinder: DataBinder{
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
							dir := tv.CurrentItem().(*windows.Directory)
							log.Println("path now :", dir.Path())
							if err := le.SetText(dir.Path()); err != nil {
								log.Fatal(err)
							}
							windows.ClearWidgets(sv)
							windows.AddImageViewWidgets(dir.Path(), sv)
						},
					},
					VSplitter{
						StretchFactor: 5,
						AssignTo:      &vs,
						Children: []Widget{
							ScrollView{
								AssignTo: &sv,
								Name:     "Pictures",
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
								DataMember: "tt",
								Buttons: []RadioButton{
									{
										Name:  "'Markdown' Text",
										Text:  "Markdown",
										Value: "a",
									},
									{
										Name:  "'HTML' Text",
										Text:  "HTML",
										Value: "b",
									},
									{
										Name:  "'URL' Text",
										Text:  "URL",
										Value: "c",
									},
									{
										Name:  "'FilePath' Text",
										Text:  "FilePath",
										Value: "d",
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
