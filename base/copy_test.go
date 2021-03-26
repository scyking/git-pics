package base_test

import (
	"gpics/base"
	"log"
	"testing"
)

func TestCopy(t *testing.T) {
	if err := base.Copy("test.png", base.URL); err != nil {
		log.Println(err)
	}
}
