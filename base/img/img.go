package img

import (
	"github.com/lxn/walk"
	"log"
)

func Shell32Icon(index int) *walk.Icon {

	ic, err := walk.NewIconFromSysDLL("shell32", index)
	if err != nil {
		log.Println(err)
	}
	return ic
}
