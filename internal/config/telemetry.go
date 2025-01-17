package config

// TelemetryConfig holds configuration for OpenTelemetry SDK
type TelemetryConfig struct {
	// Endpoint is the OTLP endpoint URL (e.g., "localhost:4317")
	Endpoint string `yaml:"endpoint"`
	// Headers are the headers to be added to OTLP requests
	Headers map[string]string `yaml:"headers"`
	// Traces configuration
	Traces TracesConfig `yaml:"traces"`
	// Metrics configuration
	Metrics MetricsConfig `yaml:"metrics"`
	// Logs configuration
	Logs LogsConfig `yaml:"logs"`
}

// TracesConfig holds configuration for trace collection
type TracesConfig struct {
	// Enabled indicates if tracing is enabled
	Enabled bool `yaml:"enabled"`
	// SampleRatio is the ratio of traces to sample (0.0 to 1.0)
	SampleRatio float64 `yaml:"sampleRatio"`
}

// MetricsConfig holds configuration for metrics collection
type MetricsConfig struct {
	// Enabled indicates if metrics collection is enabled
	Enabled bool `yaml:"enabled"`
	// ReportingInterval is how often to report metrics (in seconds)
	ReportingInterval int `yaml:"reportingInterval"`
}

// LogsConfig holds configuration for logging
type LogsConfig struct {
	// Enabled indicates if metrics collection is enabled
	Enabled bool `yaml:"enabled"`
	// Level sets the minimum enabled logging level (debug, info, warn, error)
	Level string `yaml:"level"`
}
