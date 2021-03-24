package main

import (
	"gpics/windows"
	"log"
)

func main() {

	win, err := windows.Build()

	if err != nil {
		log.Println(err)
	}
	win.Run()
}
