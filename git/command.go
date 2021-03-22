package git

import (
	"errors"
	"gpics/config"
	"os/exec"
)

func runGitCommand(name string, arg ...string) (string, error) {
	dir := config.CmdDir
	if dir == "" {
		return "", errors.New("cmd dir is error")
	}

	cmd := exec.Command(name, arg...)
	cmd.Dir = dir
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
