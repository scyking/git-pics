package files_test

import (
	"git-pics/files"
	"log"
	"testing"
)

func TestDirFiles(t *testing.T) {
	names := files.ImageFileNames("C:\\workspace\\test")
	log.Println(names)
}
