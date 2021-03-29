package git

import (
	"errors"
	"github.com/lxn/walk"
	"gpics/config"
	"log"
	"net/url"
	"strings"
	"sync"
	"time"
)

var mu = new(sync.Mutex)

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

func Version() error {
	return version("")
}

func AutoCommit() (e error) {
	dir := walk.Resources.RootDirPath()
	if err := add(dir, "."); err != nil {
		return err
	}
	if err := commit(dir, "添加图片"); err != nil {
		return err
	}
	go remoteCommit(mu)
	return nil
}

// 因网络等原因 很容易超时失败
func remoteCommit(mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()

	time.Sleep(5 * 1e9)
	if err := Pull(); err != nil {
		log.Println("pull err:", err)
		return
	}
	if err := Push(); err != nil {
		log.Println("push err:", err)
		return
	}
	log.Println("提交 成功")
}
