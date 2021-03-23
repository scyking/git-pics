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
	var tv *walk.TreeView
	var hs *walk.Splitter
	var vs *walk.Splitter
	var sv *walk.ScrollView
	var le *walk.LineEdit

	mw := new(MyMainWindow)
	mw.DBSource = DBSource()

	treeModel, err := NewDirectoryTreeModel()
	if err != nil {
		log.Fatal(err)
	}

	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    config.PName,
		MinSize:  Size{600, 400},
		Layout:   HBox{MarginsZero: true},
		DataBinder: DataBinder{
			AutoSubmit: true,
			DataSource: mw.DBSource,
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
							mw.addImageViewWidgets(dir.Path(), sv)
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
												Name:      "'Markdown' Text",
												Text:      "Markdown",
												Value:     Markdown,
												OnClicked: mw.clickRadio,
											},
											{
												Name:      "'HTML' Text",
												Text:      "HTML",
												Value:     HTML,
												OnClicked: mw.clickRadio,
											},
											{
												Name:      "'URL' Text",
												Text:      "URL",
												Value:     URL,
												OnClicked: mw.clickRadio,
											},
											{
												Name:      "'FilePath' Text",
												Text:      "FilePath",
												Value:     FilePath,
												OnClicked: mw.clickRadio,
											},
										},
									},
									PushButton{
										Text: "添加图片",
										OnClicked: func() {
											name, err := OpenImage(mw.MainWindow)
											if err != nil {
												log.Fatal(err)
											}
											AddImageViewWidget(name, sv)
										},
									},
									PushButton{
										Text: "屏幕截图",
									},
									PushButton{
										Text: "手动Push",
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
									DataSource: mw.DBSource,
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
	return mw.MainWindow, nil
}
