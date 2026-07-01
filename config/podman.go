package config

import (
	"errors"
	"slices"

	"github.com/containers/podman-tui/pdcs/registry"
	"github.com/rs/zerolog/log"
	cconfig "go.podman.io/common/pkg/config"
	"go.podman.io/podman/v6/pkg/domain/entities"
)

var (
	ErrInvalidURISchemaName    = errors.New("invalid schema name")
	ErrInvalidTCPSchemaOption  = errors.New("invalid option for tcp")
	ErrInvalidUnixSchemaOption = errors.New("invalid option for unix")
	ErrFileNotUnixSocket       = errors.New("not a unix domain socket")
	ErrEmptySSHIdentity        = errors.New("empty identity field for SSH connection")
	ErrEmptyURIDestination     = errors.New("empty URI destination")
	ErrEmptyConnectionName     = errors.New("empty connection name")
	ErrConnectionNotFound      = errors.New("connection not found")
)

func newPodmanRemoteConfig() (*entities.PodmanConfig, error) {
	log.Debug().Msg("config: loading podman remote connections")

	defaultConfig, err := cconfig.New(&cconfig.Options{
		SetDefault: true,
		Modules:    nil,
	})
	if err != nil {
		return nil, err
	}

	podmanOptions := entities.PodmanConfig{ContainersConf: &cconfig.Config{}, ContainersConfDefaultsRO: defaultConfig}

	return &podmanOptions, nil
}

func (c *AppConfig) RemoteConnections() []registry.Connection {
	rconn := make([]registry.Connection, 0)

	conns, err := c.podmanOptions.ContainersConfDefaultsRO.GetAllConnections()
	if err != nil {
		log.Err(err).Msgf("config: podman remote connection")

		return nil
	}

	log.Debug().Msgf("connections: %v", conns)

	for _, conn := range conns {
		rconn = append(rconn, registry.Connection{
			Name:     conn.Name,
			URI:      conn.URI,
			Identity: conn.Identity,
			Default:  conn.Default,
		})
	}

	return rconn
}

func (c *AppConfig) RemoveConnection(name string) error {
	return cconfig.EditConnectionConfig(func(cfg *cconfig.ConnectionsFile) error {
		delete(cfg.Connection.Connections, name)

		if cfg.Connection.Default == name {
			cfg.Connection.Default = ""
		}

		// If there are existing farm, remove the deleted connection that might be part of a farm
		for k, v := range cfg.Farm.List {
			index := slices.Index(v, name)
			if index > -1 {
				cfg.Farm.List[k] = append(v[:index], v[index+1:]...)
			}
		}

		return nil
	})
}

func (c *AppConfig) AddConnection(name string, uri string, identity string) error {
	connURI, err := validateNewConnection(name, uri, identity)
	if err != nil {
		return err
	}

	dst := cconfig.Destination{
		URI:      connURI,
		Identity: identity,
	}

	return cconfig.EditConnectionConfig(func(cfg *cconfig.ConnectionsFile) error {
		if cfg.Connection.Connections == nil {
			cfg.Connection.Connections = map[string]cconfig.Destination{
				name: dst,
			}
			cfg.Connection.Default = name
		} else {
			cfg.Connection.Connections[name] = dst
		}

		return nil
	})
}

func (c *AppConfig) SetDefaultConnection(name string) error {
	return cconfig.EditConnectionConfig(func(cfg *cconfig.ConnectionsFile) error {
		if _, found := cfg.Connection.Connections[name]; !found {
			return ErrConnectionNotFound
		}

		cfg.Connection.Default = name

		return nil
	})
}

func (c *AppConfig) GetDefaultConnection() registry.Connection {
	for _, conn := range c.RemoteConnections() {
		if conn.Default {
			return registry.Connection{
				Name:     conn.Name,
				Identity: conn.Identity,
				URI:      conn.URI,
			}
		}
	}

	return registry.Connection{}
}
