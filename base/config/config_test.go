package config_test

import (
	"log"
	"testing"
)

func TestWorkspaces(t *testing.T) {
	log.Println(Workspace())
}

func TestSettings(t *testing.T) {
	s := Settings()
	ws, ok := s.Get(WorkspaceKey)
	if !ok {
		log.Println("工作空间配置不存在")
	}
	log.Println("workspace:", ws)
}

func TestSaveWorkspace(t *testing.T) {
	cf := new(Config)
	cf.Workspace = ""

	if err := Save(cf); err != nil {
		log.Println(err)
	}

	log.Println(Workspace())
}

func TestSaveConfig(t *testing.T) {
	cf := new(Config)
	cf.Workspace = ""
	Save(cf)
}
