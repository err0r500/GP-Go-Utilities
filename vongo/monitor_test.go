package vongo

import (
	"context"
	"testing"
)

func TestMonitorWrapContext(t *testing.T) {
	newRelicMonitoringConf := ConfigWithMonitoring{}
	config := newRelicMonitoringConf.AsInterface().(ConfigInterface)
	wrapperContext := config.(ConnectionContextWrapperHook)
	ctx := wrapperContext.WrapContext(context.TODO())
	if ctx == nil {
		t.Fatalf("newRelicMonitoringConf - failed")
	}
}
