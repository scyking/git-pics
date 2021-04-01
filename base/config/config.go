package config

import (
	"errors"
	"github.com/lxn/walk"
	"io/ioutil"
	"log"
)

const (
	Author = "scyking"
	PName  = "GPics"
)

const (
	GitInfoRepositoryKey = "git.info.repository"
	GitInfoServerKey     = "git.info.server"
	GitInfoUserNameKey   = "git.info.username"
	GitInfoPasswordKey   = "git.info.password"
	GitInfoTokenKey      = "git.info.token"
	WorkspaceKey         = "workspace"
	OnQuickKey           = "on-quick"
	QuickDirKey          = "quick-dir"
	AutoCommitKey        = "auto-commit"
)

type GitInfo struct {
	Repository string
	Server     string
	UserName   string
	Password   string
	Token      string
}

type Config struct {
	GitInfo
	Workspace  string
	OnQuick    bool   //开启快捷上传
	QuickDir   string //快捷上传目录
	AutoCommit bool   //自动提交到远程
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

func Value(key string) (string, error) {
	st := Settings()
	v, ok := st.Get(key)

	if !ok {
		return "", errors.New("获取配置失败,KEY:" + key)
	}
	return v, nil
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

func Save(cf *Config) error {

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

func Reset() error {
	// 重置
	return nil
}
