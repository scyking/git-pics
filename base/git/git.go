package git

import (
	"errors"
	"github.com/lxn/walk"
	"gpics/base/config"
	"log"
	"net/url"
	"strings"
	"sync"
	"time"
)

func RepName(u string) (string, error) {
	rs, err := url.Parse(u)
	if err != nil {
		return "", nil
	}

	us := strings.Split(rs.Path, "/")
	if len(us) < 3 {
		return "", errors.New("解析仓库名称失败")
	}
	return strings.TrimSuffix(us[2], ".git"), nil
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

func Branch() (string, error) {
	dir := walk.Resources.RootDirPath()
	b, err := branch(dir)
	if err != nil {
		return "", nil
	}
	sts := strings.Split(b, " ")
	if len(sts) < 2 {
		return "", errors.New("解析当前分支失败" + b)
	}
	return strings.TrimSuffix(sts[1], "\n"), nil
}

var mu = new(sync.Mutex)

func AutoCommit() (e error) {
	ws, ok := config.Workspace()
	if !ok {
		return errors.New("自动提交失败，原因：获取工作空间失败")
	}

	if err := add(ws, "."); err != nil {
		return err
	}
	if err := commit(ws, "自动提交"); err != nil {
		return err
	}

	if v, _ := config.BoolValue(config.AutoCommitKey); v {
		if err := RemoteCommit(); err != nil {
			return err
		}
	}
	return nil
}

func RemoteCommit() error {
	ch := make(chan error)
	defer close(ch)
	go remoteCommit(mu, ch)

	select {
	case err := <-ch:
		return err
	case <-time.After(time.Second * 3):
		return errors.New("请求超时！")
	}
}

// 因网络等原因 很容易超时失败
func remoteCommit(mu *sync.Mutex, ch chan error) {
	mu.Lock()
	defer mu.Unlock()

	if err := Pull(); err != nil {
		log.Println("pull err:", err)
		ch <- errors.New("pull 失败")
		return
	}
	if err := Push(); err != nil {
		log.Println("push err:", err)
		ch <- errors.New("push 失败")
		return
	}
	ch <- nil
	log.Println("提交 成功")
}
