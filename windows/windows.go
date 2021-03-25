package windows

import (
	"gpics/base"
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

func (mw *MyMainWindow) clickRadio() {
	log.Println("textType:", mw.DBSource[base.DBTextType])
	if mw.ImageName != "" {
		if err := base.Copy(mw.ImageName, mw.DBSource[base.DBTextType]); err != nil {
			log.Fatal(err)
		}
	}
}
