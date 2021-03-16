package main

import (
	. "git-pics/windows"
	"log"
)

func main() {
	log.Println("开始启动")
	Build().Run()
	log.Println("启动成功！")
}
