package base

import (
	"errors"
	"fmt"
	"github.com/lxn/walk"
	"gpics/base/config"
	"gpics/base/git"
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

func markdown(name string) (string, error) {
	url, err := url(name)
	if err != nil {
		return "", err
	}
	v := fmt.Sprintf("![%s](%s)", name, url)
	return v, nil
}

func html(name string) (string, error) {
	url, err := url(name)
	if err != nil {
		return "", err
	}
	v := fmt.Sprintf("<img src=%q width=%q>", url, "50%")
	return v, nil
}

func url(name string) (string, error) {
	ws, ok := config.Workspace()
	if !ok {
		return "", errors.New("获取工作空间配置失败")
	}

	// 资源绝对地址
	abs := filepath.Join(walk.Resources.RootDirPath(), name)

	// 资源对于工作空间的相对地址
	rel, err := filepath.Rel(ws, abs)
	if err != nil {
		return "", err
	}

	server, err := config.StringValue(config.GitInfoServerKey)
	if err != nil {
		return "", err
	}

	rep, err := config.StringValue(config.GitInfoRepositoryKey)
	if err != nil {
		return "", err
	}

	branch, err := git.Branch()
	if err != nil {
		return "", err
	}

	url := "https://" + server + "/" + rep + "/blob/" + branch + "/" + strings.ReplaceAll(rel, "\\", "/")

	return url, nil
}

func filePath(name string) (string, error) {
	return filepath.Join(walk.Resources.RootDirPath(), name), nil
}

func pathByTextType(name string, textType int) (string, error) {
	switch textType {
	case Markdown:
		return markdown(name)
	case HTML:
		return html(name)
	case URL:
		return url(name)
	case FilePath:
		return filePath(name)
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
