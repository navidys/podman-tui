package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

type ShortCutsConfig struct {
	Views Views `json:"views"`
}

type Views struct {
	System     System     `json:"system"`
	Pods       Pods       `json:"pods"`
	Containers Containers `json:"containers"`
	Volumes    Volumes    `json:"volumes"`
	Images     Images     `json:"images"`
	Networks   Networks   `json:"networks"`
	Secrets    Secrets    `json:"secrets"`
}

type System struct {
	AddConnection    string `json:"add_connection"`
	Connect          string `json:"connect"`
	Disconnect       string `json:"disconnect"`
	DiskUsage        string `json:"disk_usage"`
	Events           string `json:"events"`
	Info             string `json:"info"`
	Prune            string `json:"prune"`
	RemoveConnection string `json:"remove_connection"`
	SetDefault       string `json:"set default"`
}
type Pods struct {
	Create  string `json:"create"`
	Inspect string `json:"inspect"`
	Kill    string `json:"kill"`
	Pause   string `json:"pause"`
	Prune   string `json:"prune"`
	Restart string `json:"restart"`
	Remove  string `json:"rm"`
	Start   string `json:"start"`
	Stats   string `json:"stats"`
	Stop    string `json:"stop"`
	Top     string `json:"top"`
	Unpause string `json:"unpause"`
}
type Containers struct {
	Attach      string `json:"attach"`
	Checkpoint  string `json:"checkpoint"`
	Commit      string `json:"commit"`
	Create      string `json:"create"`
	Diff        string `json:"diff"`
	Exec        string `json:"exec"`
	Healthcheck string `json:"healthcheck"`
	Inspect     string `json:"inspect"`
	Kill        string `json:"kill"`
	Logs        string `json:"logs"`
	Pause       string `json:"pause"`
	Port        string `json:"port"`
	Prune       string `json:"prune"`
	Rename      string `json:"rename"`
	Restore     string `json:"restore"`
	Remove      string `json:"rm"`
	Run         string `json:"run"`
	Start       string `json:"start"`
	Stats       string `json:"stats"`
	Stop        string `json:"stop"`
	Top         string `json:"top"`
	Unpause     string `json:"unpause"`
}
type Volumes struct {
	Create  string `json:"create"`
	Export  string `json:"export"`
	Import  string `json:"import"`
	Inspect string `json:"inspect"`
	Prune   string `json:"prune"`
	Remove  string `json:"remove"`
}
type Images struct {
	Build   string `json:"build"`
	Diff    string `json:"diff"`
	History string `json:"history"`
	Import  string `json:"import"`
	Inspect string `json:"inspect"`
	Prune   string `json:"prune"`
	Remove  string `json:"rm"`
	Save    string `json:"save"`
	Search  string `json:"search"`
	Tag     string `json:"tag"`
	Tree    string `json:"tree"`
	Untag   string `json:"untag"`
}
type Networks struct {
	Connect    string `json:"connect"`
	Create     string `json:"create"`
	Disconnect string `json:"disconnect"`
	Inspect    string `json:"inspect"`
	Prune      string `json:"prune"`
	Remove     string `json:"remove"`
}
type Secrets struct {
	Create  string `json:"create"`
	Inspect string `json:"inspect"`
	Remove  string `json:"rm"`
}

func NewShotcutsConfig() (*ShortCutsConfig, error) {
	var config ShortCutsConfig

	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Error().Msgf("config: failed to resolve user home directory")

		return nil, err
	}

	configPath := filepath.Join(homedir, userAppConfig)

	log.Debug().Msgf("config: loading tui configuration: %s", configPath)

	configFile, err := os.ReadFile(configPath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		log.Error().Msg("config: failed to read tui configuration")

		return nil, err
	}

	if err != nil && errors.Is(err, os.ErrNotExist) {
		// TODO - load randomly
		log.Debug().Msg("config: tui configuration not found, loading defaults")

		return &config, nil
	}

	err = json.Unmarshal(configFile, &config)
	if err != nil {
		log.Err(err).Msgf("config: failed to Unmarshal tui configuration data")
	}

	// TODO - validate data

	log.Debug().Msgf("config: shotcuts: %v", config)

	return &config, nil
}
