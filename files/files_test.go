package files_test

import (
	"gpics/files"
	"log"
	"testing"
)

func TestDirFiles(t *testing.T) {
	names, err := files.ImageFileNames("C:\\workspace\\test")
	if err != nil {
		log.Println(err)
	}

	log.Println(names)
}

func TestCopyFile(t *testing.T) {
	if err := files.CopyFile("C:\\workspace\\test\\test.png", "C:\\workspace\\test"); err != nil {
		log.Println(err)
	}
}
