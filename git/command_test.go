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
	err := Push()
	log.Println("push test：", err)
}

func TestPull(t *testing.T) {
	err := Pull()
	log.Println("pull test：", err)
}
