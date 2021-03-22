package main

import (
	"github.com/lxn/walk"
	"gpics/config"
	"gpics/windows"
	"log"
)

func main() {
	app := walk.App()
	app.SetOrganizationName(config.Author)
	app.SetProductName(config.PName)

	settings := walk.NewIniFileSettings("settings.ini")
	log.Println("setting file path：", settings.FilePath())

	if err := settings.Load(); err != nil {
		log.Fatal(err)
	}

	app.SetSettings(settings)

	win, err := windows.Build()
	if err != nil {
		log.Fatal(err)
	}
	win.Run()
}
