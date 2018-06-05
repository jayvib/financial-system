package config_test

import (
	"testing"
	. "financial-system/config"
	"os"
	"log"
	"github.com/magiconair/properties/assert"
)

func TestLoadConfig(t *testing.T) {
	path := `C:\Users\jayson.vibandor\Documents\projects\go\src\financial-system\config.json`
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	conf, err := LoadConfig(file)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, conf.Path.ClientSecret, "client_path")
	assert.Equal(t, conf.Path.TokenFile, "token_path")
	assert.Equal(t, conf.Path.SheetIDs, "sheet_path")
}

