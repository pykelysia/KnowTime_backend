package database

import (
	"knowtime/config"
	"testing"

	"github.com/pykelysia/pyketools"
)

func TestDatabaseInit(t *testing.T) {
	config.LoadEnv("../.env")
	err := InitDatabase()
	if err != nil {
		pyketools.Fatalf("%v", err)
	}
}
