package vongo

import (
	"github.com/VoodooTeam/GP-Go-Utilities/logger"
	"go.mongodb.org/mongo-driver/event"
)

// ConfigInterface interface
type ConfigInterface interface {
	GetURI() string
	GetDatabase() string
	GetMonitor() *event.CommandMonitor
}

// Config struct
type Config struct {
	URI      string
	Database string
	Monitor  *event.CommandMonitor
}

// GetURI method
func (config *Config) GetURI() string {
	return config.URI
}

// GetDatabase method
func (config *Config) GetDatabase() string {
	return config.Database
}

// GetMonitor method
func (config *Config) GetMonitor() *event.CommandMonitor {
	return config.Monitor
}

// AsInterface method
func (config *Config) AsInterface() interface{} {
	logger.Errorf("%v", interfaceName(config))
	return config
}
