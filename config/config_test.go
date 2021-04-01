package config_test

import (
	"gpics/config"
	"log"
	"testing"
)

func TestWorkspaces(t *testing.T) {
	log.Println(config.Workspace())
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
	cf := new(config.Config)
	cf.Workspace = ""

	if err := config.Save(cf); err != nil {
		log.Println(err)
	}

	log.Println(config.Workspace())
}

func TestSaveConfig(t *testing.T) {
	cf := new(config.Config)
	cf.Workspace = ""
	config.Save(cf)
}
