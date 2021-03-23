package windows

import (
	"gpics/config"
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

	names, err := files.ImageFileNames(path)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("image names:", names)

	builder := NewBuilder(parent)

	for _, name := range names {
		iv := addImageView(name, parent)
		if err := iv.Create(builder); err != nil {
			log.Fatal(err)
		}
	}

}

func AddImageViewWidget(name string, parent walk.Container) {
	builder := NewBuilder(parent)
	iv := addImageView(name, parent)
	if err := iv.Create(builder); err != nil {
		log.Fatal(err)
	}
}

func addImageView(name string, parent walk.Container) ImageView {
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
		},
	}
	return iv
}

func OpenImage(mw *walk.MainWindow) (string, error) {
	rootPath := walk.Resources.RootDirPath()
	for _, path := range config.Workspaces() {
		if path == rootPath {
			// todo 修改为提示
			log.Println("图片上传至工作空间根目录，图片无法使用！")
		}
	}

	dlg := new(walk.FileDialog)

	dlg.FilePath = rootPath
	dlg.Filter = "Image Files (*.emf;*.bmp;*.exif;*.gif;*.jpeg;*.jpg;*.png;*.tiff)|*.emf;*.bmp;*.exif;*.gif;*.jpeg;*.jpg;*.png;*.tiff"
	dlg.Title = "Select an Image"

	if ok, err := dlg.ShowOpen(mw); err != nil {
		return "", err
	} else if !ok {
		return "", nil
	}

	return files.CopyFile(dlg.FilePath, rootPath)
}
