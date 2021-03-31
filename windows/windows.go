package windows

import (
	"gpics/base"
	"gpics/config"
	"log"
)

import (
	"github.com/lxn/walk"
)

type MyMainWindow struct {
	*walk.MainWindow
	ImageName string
	DBSource  map[string]int
}

func (mw *MyMainWindow) errMBox(err error) {
	walk.MsgBox(mw.MainWindow, "错误提示", err.Error(), walk.MsgBoxIconError)
}

func (mw *MyMainWindow) infoMBox(msg string) {
	walk.MsgBox(mw.MainWindow, "消息提示", msg, walk.MsgBoxOK)
}

func (mw *MyMainWindow) clickRadio() {
	log.Println("textType:", mw.DBSource[base.DBTextType])
	if mw.ImageName != "" {
		if err := base.Copy(mw.ImageName, mw.DBSource[base.DBTextType]); err != nil {
			mw.errMBox(err)
		}
	}
}

func (mw *MyMainWindow) openDir() (string, error) {
	dlg := new(walk.FileDialog)

	ws, err := config.Workspaces()
	if err != nil {
		return "", err
	}
	log.Println("当前工作空间", ws)

	dlg.FilePath = ws
	dlg.Title = "选择工作空间"

	ok, err := dlg.ShowBrowseFolder(mw)

	if err != nil {
		return "", err
	}

	if !ok {
		return "", err
	}

	return dlg.FilePath, nil
}
