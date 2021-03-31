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

	if _, err := windows.Build().Run(); err != nil {
		log.Println(err)
	}

}
