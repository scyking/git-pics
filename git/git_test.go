package git_test

import (
	"gpics/git"
	"log"
	"testing"
)

func TestUrl(t *testing.T) {
	url, err := git.Url("")
	log.Println(url, err)
}

func TestUrlStr(t *testing.T) {
	url, err := git.UrlStr("")
	log.Println(url, err)
}
