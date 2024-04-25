// Package setup service configuration
package setup

import (
	"github.com/trevatk/go-pkg/domain"
)

// Config service config implmentation
type Config struct {
	Logger        domain.Logger    `hcl:"logger,block"`
	Server        domain.Server    `hcl:"server,block"`
	Raft          domain.Raft      `hcl:"raft,block"`
	Chain         domain.Chain     `hcl:"chain,block"`
	MessageBroker domain.Messenger `hcl:"message_broker,block"`
}

// interface compliance
var _ domain.Config = (*Config)(nil)

// New config constructor
func New() *Config {
	return &Config{
		Logger: domain.Logger{},
		Server: domain.Server{},
		Raft:   domain.Raft{},
	}
}

// GetLogger getter logger configuration
func (cfg *Config) GetLogger() domain.Logger {
	return cfg.Logger
}

// GetServer getter server configuration
func (cfg *Config) GetServer() domain.Server {
	return cfg.Server
}

// GetRaft getter raft configuration
func (cfg *Config) GetRaft() domain.Raft {
	return cfg.Raft
}

// GetChain getter chain configuration
func (cfg *Config) GetChain() domain.Chain {
	return cfg.Chain
}

// GetMessenger getter message broker configuration
func (cfg *Config) GetMessenger() domain.Messenger {
	return cfg.MessageBroker
}