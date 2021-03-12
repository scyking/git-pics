package git

import "os/exec"

// 执行任意Git命令的封装
func RunGitCommand(name string, arg ...string) (string, error) {
	gitpath := config.Config.Gitpath // 从配置文件中获取当前git仓库的路径

	cmd := exec.Command(name, arg...)
	cmd.Dir = gitpath // 指定工作目录为git仓库目录
	//cmd.Stderr = os.Stderr
	msg, err := cmd.CombinedOutput() // 混合输出stdout+stderr
	cmd.Run()

	// 报错时 exit status 1
	return string(msg), err
}

func Pull() {
	//
}

func Push() {
	//
}

func Clone() {
	//
}

func Version() {
	//
}
