package config

import (
	"fmt"
	"github.com/lxn/walk"
	"io/ioutil"
	"log"
	"strconv"
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
	TimeOutKey           = "remote-commit-timeout"
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
	TimeOut    int    //超时时间（单位：s）
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

	if _, ok := settings.Get(OnQuickKey); !ok {
		if err := settings.Put(OnQuickKey, strconv.FormatBool(false)); err != nil {
			log.Fatal(err)
		}
	}

	if _, ok := settings.Get(AutoCommitKey); !ok {
		if err := settings.Put(AutoCommitKey, strconv.FormatBool(false)); err != nil {
			log.Fatal(err)
		}
	}

	if _, ok := settings.Get(TimeOutKey); !ok {
		if err := settings.Put(TimeOutKey, "3"); err != nil {
			log.Fatal(err)
		}
	}

	if err := settings.Save(); err != nil {
		log.Fatal(err)
	}

	app.SetSettings(settings)
}

func NewConfig() *Config {
	cf := new(Config)

	cf.Workspace, _ = StringValue(WorkspaceKey)

	cf.AutoCommit, _ = BoolValue(AutoCommitKey)
	cf.OnQuick, _ = BoolValue(OnQuickKey)

	cf.QuickDir, _ = StringValue(QuickDirKey)
	cf.Repository, _ = StringValue(GitInfoRepositoryKey)
	cf.Server, _ = StringValue(GitInfoServerKey)
	cf.UserName, _ = StringValue(GitInfoUserNameKey)
	cf.Password, _ = StringValue(GitInfoPasswordKey)
	cf.Token, _ = StringValue(GitInfoTokenKey)
	cf.TimeOut, _ = IntValue(TimeOutKey)

	return cf
}

func Settings() walk.Settings {
	return walk.App().Settings()
}

func StringValue(key string) (string, error) {
	v, ok := Settings().Get(key)

	if !ok {
		return "", fmt.Errorf("获取配置失败，key=%q", key)
	}
	return v, nil
}

func IntValue(key string) (int, error) {
	// 默认超时时间
	dto := 3

	v, err := StringValue(key)
	if err != nil {
		return dto, err
	}
	r, err := strconv.ParseUint(v, 10, 0)
	if err != nil {
		return dto, err
	}
	return int(r), nil
}

func BoolValue(key string) (bool, error) {
	v, err := StringValue(key)
	if err != nil {
		return false, err
	}

	r, err := strconv.ParseBool(v)
	if err != nil {
		return false, err
	}
	return r, nil
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

	for _, d := range dirs {
		if d.IsDir() && d.Name() == ".git" {
			break
		}
		return fmt.Errorf("%q不是一个git项目根目录", cf.Workspace)
	}

	st := Settings()

	if err := st.Put(WorkspaceKey, cf.Workspace); err != nil {
		return err
	}
	if err := st.Put(AutoCommitKey, strconv.FormatBool(cf.AutoCommit)); err != nil {
		return err
	}
	if err := st.Put(OnQuickKey, strconv.FormatBool(cf.OnQuick)); err != nil {
		return err
	}
	if err := st.Put(QuickDirKey, cf.QuickDir); err != nil {
		return err
	}
	if err := st.Put(GitInfoRepositoryKey, cf.Repository); err != nil {
		return err
	}
	if err := st.Put(GitInfoServerKey, cf.Server); err != nil {
		return err
	}
	if err := st.Put(GitInfoUserNameKey, cf.UserName); err != nil {
		return err
	}
	if err := st.Put(GitInfoPasswordKey, cf.Password); err != nil {
		return err
	}
	if err := st.Put(GitInfoTokenKey, cf.Token); err != nil {
		return err
	}
	if err := st.Put(TimeOutKey, strconv.FormatInt(int64(cf.TimeOut), 10)); err != nil {
		return err
	}

	return st.Save()
}

func Reset() error {
	st := Settings()
	if err := st.Put(OnQuickKey, strconv.FormatBool(false)); err != nil {
		return err
	}
	if err := st.Put(QuickDirKey, ""); err != nil {
		return err
	}
	if err := st.Put(AutoCommitKey, strconv.FormatBool(false)); err != nil {
		return err
	}
	return nil
}
