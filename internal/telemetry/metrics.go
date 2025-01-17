package telemetry

import (
	"context"
	"time"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/config"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/sdk/metric"
)

func initMetricsProvider(ctx context.Context, cfg *config.TelemetryConfig, serviceName string) (*metric.MeterProvider, error) {
	if !cfg.Metrics.Enabled {
		return metric.NewMeterProvider(), nil
	}

	opts := []otlpmetricgrpc.Option{
		otlpmetricgrpc.WithInsecure(),
		otlpmetricgrpc.WithEndpoint(cfg.Endpoint),
		otlpmetricgrpc.WithHeaders(cfg.Headers),
	}

	exporter, err := otlpmetricgrpc.New(ctx, opts...)
	if err != nil {
		return nil, err
	}

	res, err := newResource(serviceName)
	if err != nil {
		return nil, err
	}

	interval := time.Duration(cfg.Metrics.ReportingInterval) * time.Second
	if interval < time.Second {
		interval = 10 * time.Second
	}

	mp := metric.NewMeterProvider(
		metric.WithResource(res),
		metric.WithReader(metric.NewPeriodicReader(exporter,
			metric.WithInterval(interval))),
	)

	return mp, nil
}
