package windows

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
	"strings"
)

var mainWindow MainWindow
var inTE, outTE *walk.TextEdit

func Build() MainWindow {

	mainWindow = MainWindow{
		Title:   "SCREAMO",
		MinSize: Size{600, 400},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			PushButton{
				Text:      "SCREAM",
				OnClicked: PushButtonOnClicked,
			},
		},
	}
	return mainWindow
}

func PushButtonOnClicked() {
	err := outTE.SetText(strings.ToUpper(inTE.Text()))
	if err != nil {
		log.Println("SetText error:", err)
	}
}
