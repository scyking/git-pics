package config

import (
	"errors"
	"github.com/lxn/walk"
	"log"
	"os"
	"path/filepath"
)

const (
	Author        = "scyking"
	PName         = "GPics"
	defaultWSName = "GPicsWorkspace"
)

const (
	WorkspaceKey = "workspace"
)

func init() {
	app := walk.App()
	settings := walk.NewIniFileSettings("settings.ini")
	log.Println("setting file path：", settings.FilePath())
	log.Println("init root path: ", walk.Resources.RootDirPath())

	if err := settings.Load(); err != nil {
		log.Fatal(err)
	}

	if _, ok := settings.Get(WorkspaceKey); !ok {
		if err := settings.Put(WorkspaceKey, defaultWS()); err != nil {
			log.Fatal(err)
		}
	}

	if err := settings.Save(); err != nil {
		log.Fatal(err)
	}

	app.SetSettings(settings)
}

func Settings() walk.Settings {
	return walk.App().Settings()
}

// 工作空间
func Workspaces() (string, error) {
	w, ok := Settings().Get(WorkspaceKey)
	if !ok {
		return "", errors.New("工作空间配置不存在")
	}
	return w, nil
}

func defaultWS() string {
	rp := walk.Resources.RootDirPath()
	ws := filepath.Join(rp, defaultWSName)
	if err := os.Mkdir(ws, os.ModeDir); err != nil {
		log.Fatal(err)
	}
	return ws
}
