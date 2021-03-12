package main

import (
	. "git-pics/windows"
	"log"
)

func main() {
	_, err := Build().Run()
	if err != nil {
		log.Println("启动失败！", err)
	}
}
