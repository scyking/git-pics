package windows

import (
	"git-pics/files"
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

	names := files.ImageFileNames(path)
	log.Println("image names:", names)

	builder := NewBuilder(parent)

	for _, name := range names {
		var civ *walk.ImageView
		iv := ImageView{
			AssignTo: &civ,
			Name:     name,
			Image:    name,
			Margin:   5,
			MinSize:  Size{120, 120},
			MaxSize:  Size{120, 120},
			Mode:     ImageViewModeZoom,
			OnBoundsChanged: func() {
				//civ.Background().Dispose()
			},
			OnMouseMove: func(x, y int, button walk.MouseButton) {
				/*brush, err := walk.NewSolidColorBrush(walk.RGB(159, 215, 255))
				if err != nil {
					log.Fatal(err)
				}
				civ.SetBackground(brush)*/
			},
			OnMouseDown: func(x, y int, button walk.MouseButton) {
				if button == walk.LeftButton {
					ClearImageViewBackground(parent)
					brush, err := walk.NewSolidColorBrush(walk.RGB(143, 199, 239))
					if err != nil {
						log.Fatal(err)
					}
					civ.SetBackground(brush)
				}
				if button == walk.RightButton {
					civ.SetEnabled(true)
				}
			},
		}

		if err := iv.Create(builder); err != nil {
			log.Fatal(err)
		}
	}

}
