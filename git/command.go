package git

import (
	"log"
	"os/exec"
)

func runGitCommand(dir string, arg ...string) error {

	cmd := exec.Command("git", arg...)
	cmd.Dir = dir

	log.Println("cmd:", cmd.String())

	return cmd.Run()
}

func outGitCommand(dir string, arg ...string) (string, error) {
	cmd := exec.Command("git", arg...)
	cmd.Dir = dir
	b, err := cmd.Output()
	return string(b), err
}

func add(dir string, file string) error {
	return runGitCommand(dir, "add", file)
}

func commit(dir string, msg string) error {
	return runGitCommand(dir, "commit", "-m", msg)
}

func pull(dir string) error {
	return runGitCommand(dir, "pull")
}

func push(dir string) error {
	return runGitCommand(dir, "push")
}

func clone(dir string, url string) error {
	return runGitCommand(dir, "clone", url)
}

func version(dir string) error {
	return runGitCommand(dir, "version")
}

func remote(dir string) (string, error) {
	return outGitCommand(dir, "remote", "-v")
}
