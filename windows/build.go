package windows

import (
	"gpics/base"
	"gpics/config"
	"gpics/img"
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

	ics := tbIcons()

	treeModel, err := NewDirectoryTreeModel()
	if err != nil {
		return nil, err
	}

	if err := (MainWindow{
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
			ButtonStyle: ToolBarButtonImageBeforeText,
			Items: []MenuItem{
				Action{
					Image: ics[0],
					Text:  "Clone",
				},
				Separator{},
				Action{
					Image: ics[1],
					Text:  "Pull",
				},
				Separator{},
				Action{
					Image: ics[2],
					Text:  "Push",
				},
				Separator{},
				Action{
					Image: ics[3],
					Text:  "添加图片",
					OnTriggered: func() {
						name, err := mw.openImage()
						if err != nil {
							mw.errMBox(err)
							return
						}
						if name == "" {
							return
						}
						mw.addImageViewWidget(name, sv)
					},
				},
				Separator{},
				Action{
					Image: ics[4],
					Text:  "截图",
				},
				Separator{},
				Action{
					Image: ics[5],
					Text:  "配置",
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
						AssignTo:      &tv,
						Model:         treeModel,
						StretchFactor: 1,
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
								AssignTo: &le,
								ReadOnly: true,
							},
							ScrollView{
								AssignTo: &sv,
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
	}.Create()); err != nil {
		return nil, err
	}

	return mw.MainWindow, nil
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

func tbIcons() []*walk.Icon {

	ics := []*walk.Icon{
		img.Shell32Icon(149), //clone
		img.Shell32Icon(46),  //pull
		img.Shell32Icon(146), //push
		img.Shell32Icon(36),  //添加图片
		img.Shell32Icon(34),  //截图
		img.Shell32Icon(69)}  //配置

	return ics
}
