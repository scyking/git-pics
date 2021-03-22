package windows

import (
	"errors"
	"fmt"
	"gpics/git"
	"path/filepath"
)

import (
	"github.com/lxn/walk"
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
	workspace := "C:/workspace/test"
	abs := filepath.Join(rootPath, name)

	rel, err := filepath.Rel(workspace, abs)
	if err != nil {
		return "", err
	}

	gitPath, err := git.GitPath(rootPath)
	if err != nil {
		return "", err
	}

	url := git.HTTPS + filepath.Join(gitPath, rel)
	return url, nil
}

func filePath(name string, rootPath string) (string, error) {
	return filepath.Join(rootPath, name), nil
}

func pathByTextType(name string, textType int) (string, error) {
	rootPath := walk.Resources.RootDirPath()
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

// 数据绑定
func DBSource() map[string]int {
	// 数据绑定
	db := make(map[string]int)
	// 设置text type默认类型
	db[DBTextType] = FilePath

	return db
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
	if err := walk.Clipboard().SetText(v); err != nil {
		return err
	}
	return nil
}
