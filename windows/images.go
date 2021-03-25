package windows

import (
	"errors"
	"gpics/files"
	"log"
)

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

// 将路径中图片做为ImageView组件添加到容器中
func (mw *MyMainWindow) addImageViewWidgets(path string, parent walk.Container) {

	names, err := files.ImageFileNames(path)

	if err != nil {
		mw.errMBox(err)
	}
	log.Println("image names:", names)

	builder := NewBuilder(parent)

	for _, name := range names {
		iv := mw.addImageView(name, parent)
		if err := iv.Create(builder); err != nil {
			mw.errMBox(err)
		}
	}

}

func (mw *MyMainWindow) addImageViewWidget(name string, parent walk.Container) {
	builder := NewBuilder(parent)
	iv := mw.addImageView(name, parent)
	if err := iv.Create(builder); err != nil {
		mw.errMBox(err)
	}
}

func (mw *MyMainWindow) addImageView(name string, parent walk.Container) ImageView {
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
					mw.errMBox(err)
				}
				civ.SetBackground(brush)

				textType := parent.DataBinder().DataSource().(map[string]int)[DBTextType]

				if err := Copy(civ.Name(), textType); err != nil {
					mw.errMBox(err)
				}

				mw.ImageName = civ.Name()
			}
		},
	}
	return iv
}

func (mw *MyMainWindow) openImage() (string, error) {
	rootPath := walk.Resources.RootDirPath()

	dlg := new(walk.FileDialog)

	//dlg.FilePath = rootPath
	dlg.Filter = "Image Files (*.emf;*.bmp;*.exif;*.gif;*.jpeg;*.jpg;*.png;*.tiff)|*.emf;*.bmp;*.exif;*.gif;*.jpeg;*.jpg;*.png;*.tiff"
	dlg.Title = "Select an Image"

	ok, err := dlg.ShowOpen(mw)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", errors.New("未选择文件")
	}

	return files.CopyFile(dlg.FilePath, rootPath)
}
