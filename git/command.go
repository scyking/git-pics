package git

import (
	"os/exec"
)

func runGitCommand(dir string, name string, arg ...string) error {

	cmd := exec.Command(name, arg...)
	cmd.Dir = dir

	return cmd.Run()
}

func outGitCommand(dir string, name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	cmd.Dir = dir
	b, err := cmd.Output()
	return string(b), err
}

func pull(dir string) error {
	return runGitCommand(dir, "git", "pull")
}

func push(dir string) error {
	return runGitCommand(dir, "git", "push")
}

func clone(dir string, url string) error {
	return runGitCommand(dir, "git", "clone", url)
}

func version(dir string) error {
	return runGitCommand(dir, "git", "version")
}

func remote(dir string) (string, error) {
	return outGitCommand(dir, "git", "remote", "-v")
}
