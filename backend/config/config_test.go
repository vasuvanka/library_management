package config_test

import (
	"testing"

	"github.com/vasuvanka/library_management/backend/config"
)

func TestNewConfig(t *testing.T) {
	got := config.NewConfig()

    if got == nil {
        t.Errorf("got nil but want %q", got)
    }
}

func TestNewConfigProps(t *testing.T) {
	got := config.NewConfig()

    if got == nil {
        t.Errorf("got nil but want %q", got)
    }

    if got.DbURI != "" {
        t.Errorf("got %s but want ''", got.DbURI)
    }

    if got.Port != "" {
        t.Errorf("got %s but want ''", got.Port)
    }
}

func TestConfigInit(t *testing.T){
    conf := config.NewConfig()

    err := conf.Init()

    if err != nil {
        t.Errorf("config should be defined but got %s", err.Error())
    }

    if conf.Port == "" {
        t.Errorf("got %s but want value", conf.Port)
    }
    if conf.DbURI == "" {
        t.Errorf("got %s but want value", conf.DbURI)
    }
}

func TestConfigInitDefaultValues(t *testing.T){
    conf := config.NewConfig()

    err := conf.Init()

    if err != nil {
        t.Errorf("config should be defined but got %s", err.Error())
    }

    if conf.Port != "3000" {
        t.Errorf("got %s but want 3000", conf.Port)
    }

    if conf.DbURI != "mongodb://localhost:27017/library" {
        t.Errorf("got %s but want mongodb://localhost:27017/library", conf.DbURI)
    }

}