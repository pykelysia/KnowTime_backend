package config

import (
	"testing"

	"github.com/pykelysia/pyketools"
)

func TestLoadEnv(t *testing.T) {
	LoadEnv("../.env")
	pyketools.Infof("ChatModelName: %v;", ChatModelName)
	pyketools.Infof("ImageModelName: %v;", ImageModelName)
	pyketools.Infof("BaseURL: %v;", BaseURL)
	pyketools.Infof("APIKey: %v;", APIKey)
}
