// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package batchprocessor // import "github.com/open-telemetry/otel-arrow/collector/processor/batchprocessor

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	// "go.opentelemetry.io/collector/featuregate"
	"go.opentelemetry.io/collector/processor"
)

const (
	// The value of "type" key in configuration.
	typeStr = "batch"

	defaultSendBatchSize = uint32(8192)
	defaultTimeout       = 200 * time.Millisecond
	defaultMaxBytes      = 262144

	// defaultMetadataCardinalityLimit should be set to the number
	// of metadata configurations the user expects to submit to
	// the collector.
	defaultMetadataCardinalityLimit = 1000
)

// var UseOtelForInternalMetricsfeatureGate = featuregate.GlobalRegistry().MustRegister(
// 	"telemetry.useOtelForInternalMetrics",
// 	featuregate.StageAlpha,
// 	featuregate.WithRegisterDescription("controls whether the collector uses OpenTelemetry for internal metrics"),
// )

// NewFactory returns a new factory for the Batch processor.
func NewFactory() processor.Factory {
	return processor.NewFactory(
		typeStr,
		createDefaultConfig,
		processor.WithTraces(createTraces, component.StabilityLevelStable),
		processor.WithMetrics(createMetrics, component.StabilityLevelStable),
		processor.WithLogs(createLogs, component.StabilityLevelStable))
}

func createDefaultConfig() component.Config {
	return &Config{
		SendBatchSize:            defaultSendBatchSize,
		MaxInFlightBytes:         defaultMaxBytes,
		Timeout:                  defaultTimeout,
		MetadataCardinalityLimit: defaultMetadataCardinalityLimit,
	}
}

func createTraces(
	_ context.Context,
	set processor.CreateSettings,
	cfg component.Config,
	nextConsumer consumer.Traces,
) (processor.Traces, error) {
	return newBatchTracesProcessor(set, nextConsumer, cfg.(*Config), true) 
}

func createMetrics(
	_ context.Context,
	set processor.CreateSettings,
	cfg component.Config,
	nextConsumer consumer.Metrics,
) (processor.Metrics, error) {
	return newBatchMetricsProcessor(set, nextConsumer, cfg.(*Config), true) 
}

func createLogs(
	_ context.Context,
	set processor.CreateSettings,
	cfg component.Config,
	nextConsumer consumer.Logs,
) (processor.Logs, error) {
	return newBatchLogsProcessor(set, nextConsumer, cfg.(*Config), true) 
}