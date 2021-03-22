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
		OnDropFiles: func(files []string) {
			//todo 上传文件
			log.Println(files)
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
						AssignTo:      &vs,
						StretchFactor: 5,
						Children: []Widget{
							HSplitter{
								Children: []Widget{
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
							LineEdit{
								AssignTo: &le,
								ReadOnly: true,
								Text:     "test",
							},
							ScrollView{
								AssignTo:      &sv,
								Name:          "Pictures",
								VerticalFixed: true,
								DataBinder: DataBinder{
									DataSource: db,
								},
								Layout: Flow{
									Alignment: AlignHNearVCenter,
								},
								Children: []Widget{},
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
