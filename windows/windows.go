package windows

import (
	"errors"
	"gpics/base"
	"gpics/config"
	"gpics/git"
	"log"
)

import (
	"github.com/lxn/walk"
)

type MyMainWindow struct {
	*walk.MainWindow
	tv        *MyTreeView
	sv        *walk.ScrollView
	le        *walk.LineEdit
	ImageName string
	DBSource  map[string]int
}

func (mw *MyMainWindow) errMBox(err error) {
	log.Println(err)
	walk.MsgBox(mw.MainWindow, "错误提示", err.Error(), walk.MsgBoxIconError)
}

func (mw *MyMainWindow) infoMBox(msg string) {
	walk.MsgBox(mw.MainWindow, "消息提示", msg, walk.MsgBoxOK)
}

func (mw *MyMainWindow) dropFiles(fps []string) {
	rootPath := walk.Resources.RootDirPath()
	for _, fp := range fps {
		name, err := base.CopyFile(fp, rootPath)
		if err != nil {
			mw.errMBox(err)
		} else {
			mw.addImageViewWidget(name, mw.sv)
			if err := git.AutoCommit(); err != nil {
				mw.errMBox(err)
			}
		}
	}
}

func (mw *MyMainWindow) clone() {
	var u string

	cmd, err := RunCloneDialog(mw, &u)

	if err != nil {
		mw.errMBox(err)
		return
	}

	if cmd == walk.DlgCmdOK {
		log.Println("Clone URL:", u)

		if err := git.Clone(u); err != nil {
			mw.errMBox(err)
			return
		}

		name, err := git.RepName(u)

		if err != nil {
			mw.errMBox(err)
			return
		}

		model := mw.tv.Model().(*DirectoryTreeModel)
		if len(model.roots) < 1 {
			mw.errMBox(errors.New("tree view 根节点不存在"))
			return
		}

		mw.tv.AddItem(name, model.roots[0])
	}
}

func (mw *MyMainWindow) commit() {
	if err := git.AutoCommit(); err != nil {
		mw.errMBox(err)
	}
}

func (mw *MyMainWindow) addPic() {
	name, err := mw.openImage()
	if err != nil {
		mw.errMBox(err)
	}
	if name != "" {
		mw.addImageViewWidget(name, mw.sv)
		if err := git.AutoCommit(); err != nil {
			mw.errMBox(err)
		}
	}
}

func (mw *MyMainWindow) config() {
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
		return
	}

	if cmd == walk.DlgCmdOK {

		if err := config.SaveConfig(cf); err != nil {
			mw.errMBox(err)
			return
		}

		mw.ImageName = ""

		model := mw.tv.Model().(*DirectoryTreeModel)
		root := NewDirectory(cf.Workspace, nil)
		model.roots = []*Directory{root}

		if err := mw.tv.SetModel(model); err != nil {
			mw.errMBox(err)
			return
		}

		if err := mw.tv.SetCurrentItem(root); err != nil {
			mw.errMBox(err)
			return
		}
	}
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
