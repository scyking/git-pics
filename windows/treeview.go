package windows

import (
	"github.com/lxn/walk"
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
