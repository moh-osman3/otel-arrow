// Code generated by "go.opentelemetry.io/collector/cmd/builder". DO NOT EDIT.

// Program otelarrowcol is an OpenTelemetry Collector binary.
package main

import (
	"log"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/otelcol"
)

func main() {
	info := component.BuildInfo{
		Command:     "otelarrowcol",
		Description: "OpenTelemetry Protocol with Apache Arrow development collector, for testing and evaluation",
		Version:     "0.23.0",
	}

	if err := run(otelcol.CollectorSettings{BuildInfo: info, Factories: components}); err != nil {
		log.Fatal(err)
	}
}

func runInteractive(params otelcol.CollectorSettings) error {
	cmd := otelcol.NewCommand(params)
	if err := cmd.Execute(); err != nil {
		log.Fatalf("collector server run finished with error: %v", err)
	}

	return nil
}
