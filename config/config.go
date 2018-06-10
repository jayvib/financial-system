package config

import (
	"io"
	"github.com/pkg/errors"
	"encoding/json"
	"os"
)

const defaultConfigPath = `C:\Users\jayson.vibandor\Documents\projects\go\src\financial-system\cmd\finsheet\config.json`

var conf *Config

type Config struct {
	Path struct {
		TokenFile string `json:"token_file_path"`
		ClientSecret string `json:"client_secret_path"`
		SheetIDs string `json:"sheet_ids_path"`
	} `json:"path"`
}

func LoadConfig(reader io.Reader) (*Config, error) {
	c := new(Config)
	if conf != nil {
		return conf, nil
	}
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(c)
	if err != nil {
		return nil, errors.Wrap(err, "error while decoding JSON")
	}
	conf = c

	return c, nil
}

func DefaultConfig() (*Config, error) {
	file, err := os.Open(defaultConfigPath)
	if err != nil {
		return nil, err
	}
	return LoadConfig(file)
}
