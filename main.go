package main

import (
	"gpics/git"
	"gpics/windows"
	"log"
)

func main() {

	if err := git.Version(); err != nil {
		log.Println("请检查git是否正确安装")
	}

	mw := windows.Build()

	if err := mw.Create(); err != nil {
		log.Println(err)
	}

	if _, err := mw.Run(); err != nil {
		log.Println(err)
	}

}
