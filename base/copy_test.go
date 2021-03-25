package base_test

import (
	"gpics/base"
	"log"
	"testing"
)

func TestCopy(t *testing.T) {
	if err := base.Copy("test", base.FilePath); err != nil {
		log.Println(err)
	}
}
