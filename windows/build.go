package windows

import (
	"gpics/base"
	"gpics/base/config"
	"gpics/base/img"
	"log"
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

func RunConfigDialog(owner walk.Form, cf *config.Config) (int, error) {
	var dlg *walk.Dialog
	var le *walk.LineEdit
	var db *walk.DataBinder
	var acceptPB, cancelPB *walk.PushButton

	return Dialog{
		AssignTo:      &dlg,
		Title:         "配置",
		Icon:          img.Shell32Icon(69),
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		DataBinder: DataBinder{
			AssignTo:       &db,
			Name:           "config",
			DataSource:     cf,
			ErrorPresenter: ToolTipErrorPresenter{},
		},
		Layout:  VBox{},
		MinSize: Size{380, 200},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 3},
				Children: []Widget{
					GroupBox{
						ColumnSpan: 3,
						Title:      "Git信息",
						Layout:     Grid{Columns: 3},
						Children: []Widget{
							Label{
								Text: "Repository:",
							},
							LineEdit{
								ColumnSpan: 2,
							},
							Label{
								Text: "Server:",
							},
							LineEdit{
								ColumnSpan: 2,
							},
							Label{
								Text: "UserName:",
							},
							LineEdit{
								ColumnSpan: 2,
							},
							Label{
								Text: "Password:",
							},
							LineEdit{
								ColumnSpan:   2,
								PasswordMode: true,
							},
							Label{
								Text: "Token:",
							},
							LineEdit{
								ColumnSpan: 2,
							},
						},
					},
					GroupBox{
						Title:      "快捷上传:",
						ColumnSpan: 3,
						Layout:     Grid{Columns: 3},
						Children: []Widget{
							Composite{
								ColumnSpan: 3,
								Layout:     HBox{},
								Children: []Widget{
									Label{
										Text: "开启状态:",
									},
									Composite{
										Layout: HBox{},
										Children: []Widget{
											RadioButtonGroup{
												Buttons: []RadioButton{
													{
														Text: "ON",
													},
													{
														Text: "OFF",
													},
												},
											},
										},
									},
								},
							},
							Composite{
								ColumnSpan: 3,
								Layout:     HBox{},
								Children: []Widget{
									Label{
										Text: "上传目录:",
									},
									LineEdit{
										AssignTo: &le,
										ReadOnly: true,
										Text:     Bind("QuickDir"),
									},
									ToolButton{
										Text:  "选择文件夹",
										Image: img.Shell32Icon(4),
										OnClicked: func() {
											ws, err := OpenDir(mw, le.Text())
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
						},
					},
					GroupBox{
						Title:  "其他配置",
						Layout: Grid{Columns: 3},
						Children: []Widget{
							Composite{
								ColumnSpan: 3,
								Layout:     HBox{},
								Children: []Widget{
									Label{
										Text: "自动提交:",
									},
									Composite{
										Layout: HBox{},
										Children: []Widget{
											RadioButtonGroup{
												Buttons: []RadioButton{
													{
														Text: "ON",
													},
													{
														Text: "OFF",
													},
												},
											},
										},
									},
								},
							},
							Composite{
								ColumnSpan: 3,
								Layout:     HBox{},
								Children: []Widget{
									Label{
										Text: "工作空间:",
									},
									LineEdit{
										AssignTo: &le,
										ReadOnly: true,
										Text:     Bind("Workspace"),
									},
									ToolButton{
										Text:  "选择文件夹",
										Image: img.Shell32Icon(4),
										OnClicked: func() {
											ws, err := OpenDir(mw, le.Text())
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
								mw.errMBox(err)
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
