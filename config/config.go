package config

import (
	"github.com/containers/podman-tui/pdcs/registry"
	"go.podman.io/podman/v6/pkg/domain/entities"
)

type AppConfig struct {
	podmanOptions *entities.PodmanConfig
	tuiOptions    *TUIConfig
}

func NewConfig() (*AppConfig, error) { //nolint:ireturn
	var cfg AppConfig

	// load podman remote connections config
	pconfig, err := newPodmanRemoteConfig()
	if err != nil {
		return nil, err
	}

	cfg.podmanOptions = pconfig

	tuiCfg, err := NewTUIConfig()
	if err != nil {
		return nil, err
	}

	cfg.tuiOptions = tuiCfg

	defaultConn := cfg.GetDefaultConnection()
	if defaultConn.URI != "" && defaultConn.Name != "" {
		registry.SetConnection(defaultConn)
	}

	return &cfg, nil
}
