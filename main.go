package main

import (
	"gpics/base/config"
	"gpics/base/git"
	"gpics/base/img"
	"gpics/windows"
	"log"
)

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {

	if err := git.Version(); err != nil {
		log.Fatal("请检查git是否正确安装！")
	}

	ws, ok := config.Workspace()

	if !ok {
		cmd, err := RunCloneDialog(nil, &ws)

		if err != nil {
			log.Fatal(err)
		}

		if cmd != walk.DlgCmdOK {
			return
		}

		if err := config.SetWorkspace(ws); err != nil {
			log.Fatal(err)
		}
	}

	if ws == "" {
		log.Fatal("工作空间配置不存在！")
	}

	if _, err := windows.Build().Run(); err != nil {
		log.Fatal(err)
	}
}

func RunCloneDialog(owner walk.Form, u *string) (int, error) {
	var dlg *walk.Dialog
	var le *walk.LineEdit

	var acceptPB, cancelPB *walk.PushButton

	dlgIc := img.Shell32Icon(149)
	fIc := img.Shell32Icon(4)

	return Dialog{
		AssignTo:      &dlg,
		Icon:          dlgIc,
		Title:         "选择文件夹",
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		MinSize:       Size{324, 200},
		Layout:        VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 3},
				Children: []Widget{
					Label{
						Text: "工作空间:",
					},
					LineEdit{
						AssignTo: &le,
						ReadOnly: true,
					},
					ToolButton{
						Text:  "选择文件夹",
						Image: fIc,
						OnClicked: func() {
							ws, err := windows.OpenDir(owner, le.Text())

							if err != nil {
								log.Fatal(err)
							}

							if err := le.SetText(ws); err != nil {
								log.Fatal(err)
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
							*u = le.Text()
							dlg.Accept()
						},
					},
					PushButton{
						AssignTo: &cancelPB,
						Text:     "Cancel",
						OnClicked: func() {
							dlg.Cancel()
						},
					},
				},
			},
		},
	}.Run(owner)
}
