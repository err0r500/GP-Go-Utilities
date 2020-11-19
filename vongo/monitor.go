package vongo

import (
	"context"

	"github.com/newrelic/go-agent/v3/integrations/nrmongo"
	"github.com/newrelic/go-agent/v3/newrelic"
	"go.mongodb.org/mongo-driver/event"
)

// ConfigWithMonitoring struct
type ConfigWithMonitoring struct {
	Config
}

// NewRelicMonitoring builder
func NewRelicMonitoring(original *event.CommandMonitor) *event.CommandMonitor {
	return nrmongo.NewCommandMonitor(original)
}

// NewRelicContext method
func NewRelicContext(ctx context.Context) context.Context {
	return newrelic.NewContext(context.TODO(), newrelic.FromContext(ctx))
}

// WrapContext method
func (config ConfigWithMonitoring) WrapContext(ctx context.Context) context.Context {
	return NewRelicContext(ctx)
}

// AsInterface method
func (config *ConfigWithMonitoring) AsInterface() interface{} {
	return config
}
