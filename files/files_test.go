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
