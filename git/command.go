package git

import (
	"gpics/config"
	"os/exec"
)

func runGitCommand(dir string, name string, arg ...string) (string, error) {
	if dir == "" {
		ws, err := config.Workspaces()
		if err != nil {
			return "", err
		}
		dir = ws
	}

	cmd := exec.Command(name, arg...)
	cmd.Dir = dir
	msg, err := cmd.CombinedOutput()
	return string(msg), err
}

func Pull(dir string, branch string) (string, error) {
	return runGitCommand(dir, "git", "pull", "origin", branch)
}

func Push(dir string, branch string) (string, error) {
	return runGitCommand(dir, "git", "push", "origin", branch)
}

func Clone(dir string, url string) (string, error) {
	return runGitCommand(dir, "git", "clone", url)
}

func Version() (string, error) {
	return runGitCommand("", "git", "version")
}
