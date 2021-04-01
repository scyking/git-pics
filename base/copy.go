package base

import (
	"errors"
	"fmt"
	"github.com/lxn/walk"
	"gpics/config"
	"gpics/git"
	"log"
	"path/filepath"
	"strings"
)

const DBTextType = "tt"

const (
	Markdown = iota
	HTML
	URL
	FilePath
)

func markdown(name string, rootPath string) (string, error) {
	url, err := url(name, rootPath)
	if err != nil {
		return "", nil
	}
	v := fmt.Sprintf("![%s](%s)", name, url)
	return v, nil
}

func html(name string, rootPath string) (string, error) {
	url, err := url(name, rootPath)
	if err != nil {
		return "", nil
	}
	v := fmt.Sprintf("<img src=%q width=%q>", url, "50%")
	return v, nil
}

func url(name string, rootPath string) (string, error) {
	ws, ok := config.Workspace()
	if !ok {
		return "", errors.New("获取工作空间配置失败")
	}
	abs := filepath.Join(rootPath, name)

	// 获取去掉后缀的git url
	gl, err := git.UrlStr(rootPath)

	if err != nil {
		return "", errors.New("git 命令执行错误")
	}

	// 获取资源地址相对工作空间地址的绝对地址
	rel, err := filepath.Rel(ws, abs)
	if err != nil {
		return "", err
	}

	url := gl + strings.ReplaceAll(rel, "\\", "/")

	return url, nil
}

func filePath(name string, rootPath string) (string, error) {
	return filepath.Join(rootPath, name), nil
}

func pathByTextType(name string, textType int) (string, error) {
	rootPath := walk.Resources.RootDirPath()
	ws, ok := config.Workspace()
	if !ok {
		return "", errors.New("获取工作空间配置失败")
	}
	if ws == rootPath {
		return "", errors.New("图片无法应用")
	}
	if rootPath == "" {
		return "", errors.New("copy: get 'root dir path' failed")
	}
	switch textType {
	case Markdown:
		return markdown(name, rootPath)
	case HTML:
		return html(name, rootPath)
	case URL:
		return url(name, rootPath)
	case FilePath:
		return filePath(name, rootPath)
	default:
		return "", fmt.Errorf("not support text type : %d", textType)
	}
}

// 根据配置生成“文件”path
func Copy(name string, textType int) error {
	if name == "" {
		return errors.New("copy: filename is not allowed to be empty")
	}

	v, err := pathByTextType(name, textType)
	if err != nil {
		return err
	}
	log.Println("copy value:", v)

	ok, err := walk.Clipboard().ContainsText()

	if err != nil {
		return err
	}

	if ok {
		if err := walk.Clipboard().Clear(); err != nil {
			return err
		}
	}

	return walk.Clipboard().SetText(v)
}
