package windows

import (
	"gpics/base"
	"gpics/config"
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
				name, err := base.CopyFile(fp, rootPath)
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
						ws, err := config.Workspaces()

						if err != nil {
							mw.errMBox(err)
							return
						}

						cf.Workspace = ws

						cmd, err := RunConfigDialog(mw, cf)
						if err != nil {
							mw.errMBox(err)
						}

						if cmd == walk.DlgCmdOK {
							log.Println("重置tree view root：", cf.Workspace)
							model := tv.Model().(*DirectoryTreeModel)
							root := NewDirectory(cf.Workspace, nil)
							model.roots = []*Directory{root}

							if err := tv.SetModel(model); err != nil {
								mw.errMBox(err)
							}

							if err := tv.SetCurrentItem(root); err != nil {
								mw.errMBox(err)
							}

							mw.ImageName = ""

							if err := config.SaveConfig(cf); err != nil {
								mw.errMBox(err)
								return
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

							path := tv.CurrentItem().(*Directory).Path()

							if err := le.SetText(path); err != nil {
								mw.errMBox(err)
								return
							}

							if err := walk.Resources.SetRootDirPath(path); err != nil {
								mw.errMBox(err)
							}
							ClearWidgets(sv)
							mw.addImageViewWidgets(sv)
						},
					},
					VSplitter{
						AssignTo:      &vs,
						StretchFactor: 5,
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
								AssignTo: &le,
								ReadOnly: true,
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

func RunConfigDialog(owner walk.Form, cf *config.Config) (int, error) {
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
			DataSource:     cf,
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
						ReadOnly: true,
						Text:     Bind("Workspace"),
						OnMouseDown: func(x, y int, button walk.MouseButton) {
							if button == walk.LeftButton {
								// todo 打开文件夹
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
