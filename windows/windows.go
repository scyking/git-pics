package windows

const DBTextType = "tt"

const (
	Markdown = iota
	HTML
	URL
	FilePath
)

func DBSource() map[string]uint8 {
	// 数据绑定
	db := make(map[string]uint8)
	// 设置text type默认类型
	db[DBTextType] = FilePath
	return db
}
