package config

const (
	Author = "scyking"
	PName  = "GPics"
)

type GitConfig struct {
	Urls      []string // git 地址
	Workspace string   // git 仓库所在工作目录
}

func Init() {

}
