// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package blockingmemorylimiterextension // import "github.com/open-telemetry/otel-arrow/collector/blockingmemorylimiterextension"

//go:generate mdatagen metadata.yaml

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)
var	Type      = component.MustNewType("blocking_memory_limiter")
const ExtensionStability = component.StabilityLevelDevelopment

// NewFactory returns a new factory for the Memory Limiter extension.
func NewFactory() extension.Factory {
	return extension.NewFactory(
		Type,
		createDefaultConfig,
		createExtension,
		ExtensionStability)
}

// CreateDefaultConfig creates the default configuration for extension. Notice
// that the default configuration is expected to fail for this extension.
func createDefaultConfig() component.Config {
	return &Config{
		MemoryLimitMiB: 32,
		Timeout: 1 * time.Nanosecond,
	}
}

func createExtension(_ context.Context, set extension.CreateSettings, cfg component.Config) (extension.Extension, error) {
	return newBlockingMLExtension(cfg.(*Config), set.TelemetrySettings.Logger)
}