package windows

import (
	"gpics/base"
	"gpics/config"
	"gpics/img"
	"log"
	"net/url"
)

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

var mw = new(MyMainWindow)

func init() {
	db := make(map[string]int)
	db[base.DBTextType] = base.FilePath
	mw.DBSource = db

	mw.tv = new(MyTreeView)
}

func tbIcons() []*walk.Icon {
	return []*walk.Icon{
		img.Shell32Icon(149), //clone
		img.Shell32Icon(46),  //pull
		img.Shell32Icon(146), //push
		img.Shell32Icon(36),  //添加图片
		img.Shell32Icon(34),  //截图
		img.Shell32Icon(69)}  //配置
}

func Build() MainWindow {

	ics := tbIcons()

	tm, err := NewDirectoryTreeModel()
	if err != nil {
		log.Fatal(err)
	}
	treeModel := tm

	return MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    config.PName,
		MinSize:  Size{800, 495},
		Layout: HBox{
			MarginsZero: true,
		},
		DataBinder: DataBinder{
			AutoSubmit: true,
			DataSource: mw.DBSource,
		},
		OnDropFiles: mw.dropFiles,
		ToolBar: ToolBar{
			ButtonStyle: ToolBarButtonImageBeforeText,
			Items: []MenuItem{
				Action{
					Image:       ics[0],
					Text:        "Clone",
					OnTriggered: mw.clone,
				},
				Separator{},
				Action{
					Image:       ics[2],
					Text:        "手动提交",
					OnTriggered: mw.commit,
				},
				Separator{},
				Action{
					Image:       ics[3],
					Text:        "添加图片",
					OnTriggered: mw.addPic,
				},
				Separator{},
				Action{
					Image:   ics[4],
					Enabled: false,
					Text:    "截图",
				},
				Separator{},
				Action{
					Image:       ics[5],
					Text:        "配置",
					OnTriggered: mw.config,
				},
			},
		},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TreeView{
						AssignTo:             &mw.tv.TreeView,
						Model:                treeModel,
						StretchFactor:        1,
						OnMouseDown:          mw.rightClick,
						OnCurrentItemChanged: mw.itemChanged,
					},
					VSplitter{
						StretchFactor: 3,
						Children: []Widget{
							HSplitter{
								Children: []Widget{
									RadioButtonGroup{
										DataMember: base.DBTextType,
										Buttons: []RadioButton{
											{
												Name:      "'Markdown' Text",
												Text:      "Markdown",
												Value:     base.Markdown,
												OnClicked: mw.clickRadio,
											},
											{
												Name:      "'HTML' Text",
												Text:      "HTML",
												Value:     base.HTML,
												OnClicked: mw.clickRadio,
											},
											{
												Name:      "'URL' Text",
												Text:      "URL",
												Value:     base.URL,
												OnClicked: mw.clickRadio,
											},
											{
												Name:      "'FilePath' Text",
												Text:      "FilePath",
												Value:     base.FilePath,
												OnClicked: mw.clickRadio,
											},
										},
									},
								},
							},
							LineEdit{
								AssignTo: &mw.le,
								ReadOnly: true,
							},
							ScrollView{
								AssignTo: &mw.sv,
								Name:     "Pictures",
								DataBinder: DataBinder{
									DataSource: mw.DBSource,
								},
								Layout: Flow{
									Alignment: AlignHNearVNear,
								},
								Children: []Widget{},
							},
						},
					},
				},
			},
		},
	}
}

////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////

func RunCreateDirDialog(owner walk.Form, u *string) (int, error) {
	var dlg *walk.Dialog
	var le *walk.LineEdit

	var acceptPB, cancelPB *walk.PushButton

	return Dialog{
		AssignTo:      &dlg,
		Title:         "新建文件夹",
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		MinSize:       Size{324, 200},
		Layout:        VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{
						Text: "文件夹名称:",
					},
					LineEdit{
						AssignTo: &le,
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					PushButton{
						AssignTo: &acceptPB,
						Text:     "OK",
						OnClicked: func() {
							*u = le.Text()
							dlg.Accept()
						},
					},
					PushButton{
						AssignTo:  &cancelPB,
						Text:      "Cancel",
						OnClicked: func() { dlg.Cancel() },
					},
				},
			},
		},
	}.Run(owner)
}

func RunCloneDialog(owner walk.Form, u *string) (int, error) {
	var dlg *walk.Dialog
	var le *walk.LineEdit

	var acceptPB, cancelPB *walk.PushButton

	return Dialog{
		AssignTo:      &dlg,
		Title:         "Clone",
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		MinSize:       Size{324, 200},
		Layout:        VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 3},
				Children: []Widget{
					Label{
						Text: "URL:",
					},
					LineEdit{
						AssignTo: &le,
					},
					PushButton{
						Text: "Test",
						OnClicked: func() {
							// 检查是否是一个url
							_, err := url.ParseRequestURI(le.Text())
							if err != nil {
								mw.errMBox(err)
								return
							}
							mw.infoMBox("测试成功")
						},
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					PushButton{
						AssignTo: &acceptPB,
						Text:     "OK",
						OnClicked: func() {
							*u = le.Text()
							dlg.Accept()
						},
					},
					PushButton{
						AssignTo:  &cancelPB,
						Text:      "Cancel",
						OnClicked: func() { dlg.Cancel() },
					},
				},
			},
		},
	}.Run(owner)

}

func RunConfigDialog(owner walk.Form, cf *config.Config) (int, error) {
	var dlg *walk.Dialog
	var le *walk.LineEdit
	var db *walk.DataBinder
	var acceptPB, cancelPB *walk.PushButton

	ic := img.Shell32Icon(4)

	return Dialog{
		AssignTo:      &dlg,
		Title:         "配置",
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		DataBinder: DataBinder{
			AssignTo:       &db,
			Name:           "config",
			DataSource:     cf,
			ErrorPresenter: ToolTipErrorPresenter{},
		},
		MinSize: Size{324, 200},
		Layout:  VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 3},
				Children: []Widget{
					Label{
						Text: "Workspace:",
					},
					LineEdit{
						AssignTo: &le,
						ReadOnly: true,
						Text:     Bind("Workspace"),
					},
					ToolButton{
						Name:  "选择",
						Image: ic,
						OnClicked: func() {
							ws, err := mw.openDir()
							if ws == "" {
								return
							}

							if err != nil {
								mw.errMBox(err)
							}

							if err := le.SetText(ws); err != nil {
								mw.errMBox(err)
								return
							}
						},
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					PushButton{
						AssignTo: &acceptPB,
						Text:     "OK",
						OnClicked: func() {
							if err := db.Submit(); err != nil {
								log.Print(err)
								return
							}

							dlg.Accept()
						},
					},
					PushButton{
						AssignTo:  &cancelPB,
						Text:      "Cancel",
						OnClicked: func() { dlg.Cancel() },
					},
				},
			},
		},
	}.Run(owner)
}
