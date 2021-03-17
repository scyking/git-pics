package windows

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

// 返回地址中图片名称数组
func ImageFileNames(filePath string) []string {
	var names []string

	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if !file.IsDir() && isImageFile(file.Name()) {
			names = append(names, file.Name())
		}
	}
	return names
}

func isImageFile(name string) bool {
	ext := filepath.Ext(name)
	// todo 匹配图片格式
	// Image Files (*.emf;*.bmp;*.exif;*.gif;*.jpeg;*.jpg;*.png;*.tiff)|*.emf;*.bmp;*.exif;*.gif;*.jpeg;*.jpg;*.png;*.tiff
	if ext == ".png" {
		return true
	}

	return false
}
