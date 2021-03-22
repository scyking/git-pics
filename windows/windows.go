package windows

const DBTextType = "tt"

const (
	Markdown = iota
	HTML
	URL
	FilePath
)

// 数据绑定
func DBSource() map[string]int {
	// 数据绑定
	db := make(map[string]int)
	// 设置text type默认类型
	db[DBTextType] = FilePath

	return db
}

// 根据配置生成“文件”path
func Copy() error {
	return nil
}
