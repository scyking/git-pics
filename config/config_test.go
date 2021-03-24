package config_test

import (
	"gpics/config"
	"log"
	"testing"
)

func TestWorkspaces(t *testing.T) {
	ws, err := config.Workspaces()
	if err != nil {
		log.Println("err:", err)
	}
	log.Println("workspace:", ws)
}

func TestSettings(t *testing.T) {
	s := config.Settings()
	ws, ok := s.Get(config.WorkspaceKey)
	if !ok {
		log.Println("工作空间配置不存在")
	}
	log.Println("workspace:", ws)
}

func TestSaveWorkspace(t *testing.T) {
	if err := config.SaveWorkspace("test"); err != nil {
		log.Println(err)
	}
	if ws, err := config.Workspaces(); err != nil {
		log.Println(err)
	} else {
		log.Println(ws)
	}
}
