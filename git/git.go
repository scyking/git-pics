package git

import (
	"errors"
	"github.com/lxn/walk"
	"gpics/config"
	"net/url"
	"strings"
)

// 获取git url
func Url(dir string) (*url.URL, error) {
	r, err := remote(dir)
	if err != nil {
		return nil, err
	}

	st := strings.Split(r, " ")
	s := strings.Split(st[0], "\t")
	return url.Parse(s[1])
}

func UrlStr(dir string) (string, error) {
	u, err := Url(dir)
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(u.String(), ".git"), nil
}

func Clone(url string) error {

	dir, err := config.Workspaces()
	if err != nil {
		return err
	}

	if url == "" {
		return errors.New("Git 地址不正确 ")
	}

	return clone(dir, url)
}

func Pull() error {
	dir := walk.Resources.RootDirPath()

	return pull(dir)
}

func Push() error {
	dir := walk.Resources.RootDirPath()

	return push(dir)
}
