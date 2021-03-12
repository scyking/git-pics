package git

import (
	"git-pics/config"
	"os/exec"
)

var gitConfig config.GitConfig

func runGitCommand(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	cmd.Dir = "C:\\"
	msg, err := cmd.CombinedOutput()
	return string(msg), err
}

func Pull(branch string) (string, error) {
	return runGitCommand("git", "pull", "origin", branch)
}

func Push(branch string) (string, error) {
	return runGitCommand("git", "push", "origin", branch)
}

func Clone(url string) (string, error) {
	return runGitCommand("git", "clone", url)
}

func Version() (string, error) {
	return runGitCommand("git", "version")
}
