package client

import (
	"encoding/json"
	"io/ioutil"
)

type AppConfig struct {
	ApiToken  string           `json:"token"`
	ProjectId string           `json:"projectId"`
	Languages []LanguageConfig `json:"languages"`
	Type      string           `json:"type"`
	Tags      []string         `json:"tags"`
}

type LanguageConfig struct {
	LangCode   string `json:"code"`
	ExportPath string `json:"path"`
}

func FromFile(path string) (*AppConfig, error) {
	var appConfig AppConfig
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(file, &appConfig); err != nil {
		return nil, err
	} else {
		return &appConfig, nil
	}
}
