package windows_test

import (
	"git-pics/windows"
	"log"
	"testing"
)

func TestDirFiles(t *testing.T) {
	names := windows.DirFiles("C:\\workspace\\test")
	log.Println(names)
}
