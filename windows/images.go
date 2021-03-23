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

func OpenImage(mw *walk.MainWindow) error {
	dlg := new(walk.FileDialog)

	dlg.FilePath = walk.Resources.RootDirPath()
	dlg.Filter = "Image Files (*.emf;*.bmp;*.exif;*.gif;*.jpeg;*.jpg;*.png;*.tiff)|*.emf;*.bmp;*.exif;*.gif;*.jpeg;*.jpg;*.png;*.tiff"
	dlg.Title = "Select an Image"

	if ok, err := dlg.ShowOpen(mw); err != nil {
		return err
	} else if !ok {
		return nil
	}

	log.Println("select image path : ", dlg.FilePath)

	return nil
}
