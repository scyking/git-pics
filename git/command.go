package git

import (
	"log"
	"os/exec"
)

func RunGitCommand(name string, arg ...string) string {
	cmd := exec.Command(name, arg...)
	msg, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal("git 命令执行失败！", err)
	}
	return string(msg)
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
	result := RunGitCommand("git", "version")
	log.Println("git version:", result)
}
