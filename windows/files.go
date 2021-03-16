package windows

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

func DirFiles(filePath string) []string {
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
	if ext == ".png" {
		return true
	}

	return false
}
