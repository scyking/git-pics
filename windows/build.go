package windows

import (
	"gpics/config"
	"gpics/files"
	"log"
)

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

var mw = new(MyMainWindow)

func init() {
	db := make(map[string]int)
	db[DBTextType] = FilePath
	mw.DBSource = db
}

func Build() (*walk.MainWindow, error) {
	var tv *walk.TreeView
	var hs *walk.Splitter
	var vs *walk.Splitter
	var sv *walk.ScrollView
	var le *walk.LineEdit

	treeModel, err := NewDirectoryTreeModel()
	if err != nil {
		return nil, err
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
		OnDropFiles: func(fps []string) {
			rootPath := walk.Resources.RootDirPath()
			for _, fp := range fps {
				name, err := files.CopyFile(fp, rootPath)
				if err != nil {
					mw.errMBox(err)
					return
				}
				mw.addImageViewWidget(name, sv)
			}
		},
		ToolBar: ToolBar{
			Font:        Font{PointSize: 15},
			ButtonStyle: ToolBarButtonTextOnly,
			Items: []MenuItem{
				Action{
					Text: "Clone",
				},
				Separator{},
				Action{
					Text: "Pull",
				},
				Separator{},
				Action{
					Text: "Push",
				},
				Separator{},
				Action{
					Text: "添加图片",
					OnTriggered: func() {
						name, err := mw.openImage()
						if err != nil {
							mw.errMBox(err)
							return
						}
						mw.addImageViewWidget(name, sv)
					},
				},
				Separator{},
				Action{
					Text: "屏幕截图",
				},
				Separator{},
				Action{
					Text: "配置",
					OnTriggered: func() {
						cf := new(config.Config)
						cf.Workspace, err = config.Workspaces()
						if err != nil {
							mw.errMBox(err)
						}
						if cmd, err := RunConfigDialog(mw, cf); err != nil {
							log.Print(err)
						} else if cmd == walk.DlgCmdOK {
							if err := config.SaveWorkspace(cf.Workspace); err != nil {
								mw.errMBox(err)
							}
						}
					},
				},
			},
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
								mw.errMBox(err)
								return
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

func RunConfigDialog(owner walk.Form, config *config.Config) (int, error) {
	var dlg *walk.Dialog
	var db *walk.DataBinder
	var acceptPB, cancelPB *walk.PushButton

	return Dialog{
		AssignTo:      &dlg,
		Title:         "配置",
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		DataBinder: DataBinder{
			AssignTo:       &db,
			Name:           "config",
			DataSource:     config,
			ErrorPresenter: ToolTipErrorPresenter{},
		},
		MinSize: Size{300, 300},
		Layout:  VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{
						Text: "Workspace:",
					},
					LineEdit{
						Text: Bind("Workspace"),
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
