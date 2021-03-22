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

	config.InitConfig()

	win, err := windows.Build()
	if err != nil {
		log.Fatal(err)
	}
	win.Run()
}
