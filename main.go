package main

import (
	"gpics/git"
	"gpics/windows"
	"log"
)

func main() {

	if err := git.Version(); err != nil {
		log.Fatal("请检查git是否正确安装")
	}

	mw, err := windows.Build()

	if err != nil {
		log.Fatal(err)
	}

	mw.Run()

}
