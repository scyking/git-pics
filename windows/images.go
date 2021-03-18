package windows

import (
	"log"
)

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

// 将路径中图片做为ImageView组件添加到容器中
func AddImageViewWidgets(path string, parent walk.Container) {

	if err := walk.Resources.SetRootDirPath(path); err != nil {
		log.Fatal(err)
	}

	names := ImageFileNames(path)
	log.Println("image names:", names)

	builder := NewBuilder(parent)

	for _, name := range names {
		iv := ImageView{
			Image:   name,
			Margin:  5,
			MinSize: Size{120, 120},
			MaxSize: Size{120, 120},
			Mode:    ImageViewModeZoom,
		}

		if err := iv.Create(builder); err != nil {
			log.Fatal(err)
		}
	}

}
