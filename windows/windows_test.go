package windows_test

import (
	"github.com/lxn/walk"
	"log"
	"testing"
)

func TestBuild(t *testing.T) {
	img, err := walk.NewImageFromFileForDPI("../img/plus.png", 96)
	if err != nil {
		log.Println(err)
	}

	ic, err := walk.NewIconFromFile("../img/open.png")
	if err != nil {
		log.Println(err)
	}

	log.Println(img)
	log.Println(ic)
}
