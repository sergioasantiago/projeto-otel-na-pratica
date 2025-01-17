package telemetry

import (
	"context"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/config"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/trace"
)

func initTraceProvider(ctx context.Context, cfg *config.TelemetryConfig, serviceName string) (*trace.TracerProvider, error) {
	if !cfg.Traces.Enabled {
		return trace.NewTracerProvider(), nil
	}

	opts := []otlptracegrpc.Option{
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(cfg.Endpoint),
		otlptracegrpc.WithHeaders(cfg.Headers),
	}

	exporter, err := otlptracegrpc.New(ctx, opts...)
	if err != nil {
		return nil, err
	}

	res, err := newResource(serviceName)
	if err != nil {
		return nil, err
	}

	samplerRatio := cfg.Traces.SampleRatio
	if samplerRatio <= 0 {
		samplerRatio = 1.0
	}

	tp := trace.NewTracerProvider(
		trace.WithResource(res),
		trace.WithBatcher(exporter),
		trace.WithSampler(trace.TraceIDRatioBased(samplerRatio)),
	)

	return tp, nil
}
