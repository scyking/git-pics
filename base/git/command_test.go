package git

import (
	"log"
	"testing"
)

func TestVersion(t *testing.T) {
	err := version("")
	log.Println(err)
}

func TestClone(t *testing.T) {
	err := clone("", "")
	log.Println("clone test：", err)
}

func TestPush(t *testing.T) {
	err := push("")
	log.Println("push test：", err)
}

func TestPull(t *testing.T) {
	err := pull("")
	log.Println("pull test：", err)
}

func TestAdd(t *testing.T) {
	err := add("", ".")
	log.Println(err)
}

func TestCommit(t *testing.T) {
	err := commit("", "test")
	log.Println(err)
}
