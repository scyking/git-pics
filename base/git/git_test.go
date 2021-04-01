package git_test

import (
	"gpics/base/git"
	"log"
	"net/url"
	"testing"
)

func TestUrl(t *testing.T) {
	log.Println(url.Parse("https://github.com/scyking"))
	log.Println(url.Parse("https://github.com//scyking"))
	log.Println(url.Parse("https://github.com/scyking/"))
}

func TestAutoCommit(t *testing.T) {
	if err := git.AutoCommit(); err != nil {
		log.Println("err:", err)
	}
	if err := git.AutoCommit(); err != nil {
		log.Println("err:", err)
	}
	if err := git.AutoCommit(); err != nil {
		log.Println("err:", err)
	}
}

func TestRepName(t *testing.T) {
	name, err := git.RepName("https://github.com/scyking/test.git")
	log.Println(name, err)
}
