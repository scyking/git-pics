package git_test

import (
	"gpics/git"
	"log"
	"testing"
)

func TestVersion(t *testing.T) {
	msg, err := git.Version()
	log.Println("version test：", msg, err)
}

func TestClone(t *testing.T) {
	msg, err := git.Clone("https://github.com/scyking/my-pics.git")
	log.Println("clone test：", msg, err)
}

func TestPush(t *testing.T) {
	msg, err := git.Push("master")
	log.Println("push test：", msg, err)
}

func TestPull(t *testing.T) {
	msg, err := git.Pull("master")
	log.Println("pull test：", msg, err)
}
