package git

import (
	"errors"
	"github.com/lxn/walk"
	"log"
	"net/url"
	"strings"
	"sync"
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
	return sts[1], nil
}

var mu = new(sync.Mutex)

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
