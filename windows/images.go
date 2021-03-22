package windows

import (
	"gpics/files"
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
			OnMouseDown: func(x, y int, button walk.MouseButton) {
				if button == walk.LeftButton {
					ClearImageViewBackground(parent)
					brush, err := walk.NewSolidColorBrush(walk.RGB(143, 199, 239))
					if err != nil {
						log.Fatal(err)
					}
					civ.SetBackground(brush)

					textType := parent.DataBinder().DataSource().(map[string]int)[DBTextType]

					if err := Copy(civ.Name(), textType); err != nil {
						log.Fatal(err)
					}
				}
				if button == walk.RightButton {

				}
			},
		}

		if err := iv.Create(builder); err != nil {
			log.Fatal(err)
		}
	}

}
