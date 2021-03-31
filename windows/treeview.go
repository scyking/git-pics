package windows

import (
	"github.com/lxn/walk"
	"os"
	"path/filepath"
)

type MyTreeView struct {
	*walk.TreeView
}

func (tv *MyTreeView) AddItem(name string, parent *Directory) {
	tv.SetSuspended(true)
	defer tv.SetSuspended(false)

	model := tv.Model().(*DirectoryTreeModel)
	nd := NewDirectory(name, parent)
	parent.children = append(parent.children, nd)
	model.PublishItemsReset(parent)
}

func (mw *MyMainWindow) itemChanged() {
	path := mw.tv.CurrentItem().(*Directory).Path()

	if err := walk.Resources.SetRootDirPath(path); err != nil {
		mw.errMBox(err)
	}

	if err := mw.le.SetText(path); err != nil {
		mw.errMBox(err)
		return
	}

	ClearWidgets(mw.sv)
	mw.addImageViewWidgets(mw.sv)
}

func (mw *MyMainWindow) rightClick(x, y int, button walk.MouseButton) {
	if button != walk.RightButton {
		return
	}
	item := mw.tv.ItemAt(x, y)
	if item == nil {
		return
	}

	if item.Parent() == nil {
		return
	}

	var fn string

	cmd, err := RunCreateDirDialog(mw, &fn)
	if err != nil {
		mw.errMBox(err)
		return
	}

	if cmd == walk.DlgCmdOK {
		path := filepath.Join(item.(*Directory).Path(), fn)

		if err := os.Mkdir(path, os.ModeDir); err != nil {
			mw.errMBox(err)
			return
		}
		mw.tv.AddItem(fn, item.(*Directory))
	}
}
