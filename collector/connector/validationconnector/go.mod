module github.com/open-telemetry/otel-arrow/collector/connector/validationconnector

go 1.21

toolchain go1.21.4

require (
	github.com/open-telemetry/otel-arrow v0.20.0
	go.opentelemetry.io/collector v0.97.0
	go.opentelemetry.io/collector/component v0.97.0
	go.opentelemetry.io/collector/connector v0.97.0
	go.opentelemetry.io/collector/consumer v0.97.0
	go.opentelemetry.io/collector/pdata v1.4.0
	go.uber.org/zap v1.27.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-viper/mapstructure/v2 v2.0.0-alpha.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.97.0 // indirect
	go.opentelemetry.io/collector/confmap v0.97.0 // indirect
	go.opentelemetry.io/otel v1.24.0 // indirect
	go.opentelemetry.io/otel/metric v1.24.0 // indirect
	go.opentelemetry.io/otel/trace v1.24.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240221002015-b0ce06bbee7c // indirect
	google.golang.org/grpc v1.62.1 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
replace github.com/open-telemetry/otel-arrow/collector/cmd/otelarrowcol => ../../../collector/cmd/otelarrowcol
replace github.com/open-telemetry/otel-arrow/collector/examples/printer => ../../../collector/examples/printer
replace github.com/open-telemetry/otel-arrow/collector/exporter/fileexporter => ../../../collector/exporter/fileexporter
replace github.com/open-telemetry/otel-arrow/collector/exporter/otelarrowexporter => ../../../collector/exporter/otelarrowexporter
replace github.com/open-telemetry/otel-arrow/collector => ../../../collector
replace github.com/open-telemetry/otel-arrow/collector/processor/concurrentbatchprocessor => ../../../collector/processor/concurrentbatchprocessor
replace github.com/open-telemetry/otel-arrow/collector/processor/experimentprocessor => ../../../collector/processor/experimentprocessor
replace github.com/open-telemetry/otel-arrow/collector/processor/obfuscationprocessor => ../../../collector/processor/obfuscationprocessor
replace github.com/open-telemetry/otel-arrow/collector/receiver/filereceiver => ../../../collector/receiver/filereceiver
replace github.com/open-telemetry/otel-arrow/collector/receiver/otelarrowreceiver => ../../../collector/receiver/otelarrowreceiver
replace github.com/open-telemetry/otel-arrow => ../../..
