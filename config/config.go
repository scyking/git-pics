package config

import (
	"errors"
	"github.com/lxn/walk"
	"io/ioutil"
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

type Config struct {
	Workspace  string
	AutoCommit bool
}

func init() {
	app := walk.App()
	app.SetOrganizationName(Author)
	app.SetProductName(PName)

	settings := walk.NewIniFileSettings("settings.ini")
	log.Println("配置文件路径：", settings.FilePath())
	log.Println("初始资源根路径: ", walk.Resources.RootDirPath())

	if err := settings.Load(); err != nil {
		log.Fatal(err)
	}

	app.SetSettings(settings)
}

func Settings() walk.Settings {
	return walk.App().Settings()
}

func Workspace() (string, bool) {
	return Settings().Get(WorkspaceKey)
}

func SetWorkspace(ws string) error {
	st := Settings()

	if err := st.Put(WorkspaceKey, ws); err != nil {
		return err
	}

	return st.Save()
}

func SaveConfig(cf *Config) error {

	dirs, err := ioutil.ReadDir(cf.Workspace)
	if err != nil {
		return nil
	}

	if len(dirs) > 0 {
		return errors.New("请选择空文件夹作为工作空间")
	}

	st := Settings()

	if err := st.Put(WorkspaceKey, cf.Workspace); err != nil {
		return err
	}

	return st.Save()
}

func defaultWS() string {
	rp := walk.Resources.RootDirPath()
	ws := filepath.Join(rp, defaultWSName)
	if err := os.Mkdir(ws, os.ModeDir); err != nil {
		log.Fatal(err)
	}
	return ws
}
