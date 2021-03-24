package windows_test

import (
	"gpics/windows"
	"log"
	"testing"
)

func TestCopy(t *testing.T) {
	if err := windows.Copy("test", windows.FilePath); err != nil {
		log.Println(err)
	}
}
