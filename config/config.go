package config

import (
	"github.com/lxn/walk"
	"log"
)

const (
	Author = "scyking"
	PName  = "GPics"
)

const (
	Workspace = "workspace"
)

// 更换成 walk.Resources.RootDirPath()
var CmdDir = ""

func InitConfig() {
	app := walk.App()
	settings := walk.NewIniFileSettings("settings.ini")
	log.Println("setting file path：", settings.FilePath())
	log.Println("init root path: ", walk.Resources.RootDirPath())

	if err := settings.Load(); err != nil {
		log.Fatal(err)
	}
	//cmd.Dir = "C:\\ideaproject\\my-pics"
	if v, ok := settings.Get(Workspace); ok {
		CmdDir = v
	}

	CmdDir = "C:\\ideaproject\\my-pics"

	app.SetSettings(settings)
}

func Settings() walk.Settings {
	app := walk.App()
	return app.Settings()
}

// 工作空间
func Workspaces() []string {
	//todo 目前不支持多工作空间
	return []string{"C:/workspace/test"}
}
