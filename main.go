package main

import (
	. "git-pics/windows"
	"log"
)

func main() {
	win, err := Build()
	if err != nil {
		log.Fatal(err)
	}
	win.Run()
}
