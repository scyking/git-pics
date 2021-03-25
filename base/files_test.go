package base_test

import (
	"gpics/base"
	"log"
	"testing"
)

func TestDirFiles(t *testing.T) {
	names, err := base.ImageFileNames("C:\\workspace\\test")
	if err != nil {
		log.Println(err)
	}

	log.Println(names)
}

func TestCopyFile(t *testing.T) {
	if _, err := base.CopyFile("C:\\workspace\\test\\test.png", "C:\\workspace\\test"); err != nil {
		log.Println(err)
	}
}
