package gpics

import (
	"github.com/lxn/walk"
	"log"
)

const (
	Author = "scyking"
	PName  = "GPics"
)

func main() {
	app := walk.App()
	app.SetOrganizationName(Author)
	app.SetProductName(PName)

	settings := walk.NewIniFileSettings("settings.ini")

	if err := settings.Load(); err != nil {
		log.Fatal(err)
	}

	app.SetSettings(settings)

	win, err := Build()
	if err != nil {
		log.Fatal(err)
	}
	win.Run()
}
