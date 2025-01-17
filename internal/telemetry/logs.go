package telemetry

import (
	"context"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/config"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/sdk/log"
)

func initLoggerProvider(ctx context.Context, cfg *config.TelemetryConfig, serviceName string) (*log.LoggerProvider, error) {
	if !cfg.Logs.Enabled {
		return log.NewLoggerProvider(), nil
	}

	opts := []otlploggrpc.Option{
		otlploggrpc.WithInsecure(),
		otlploggrpc.WithEndpoint(cfg.Endpoint),
		otlploggrpc.WithHeaders(cfg.Headers),
	}

	exporter, err := otlploggrpc.New(ctx, opts...)
	if err != nil {
		return nil, err
	}

	res, err := newResource(serviceName)
	if err != nil {
		return nil, err
	}

	lp := log.NewLoggerProvider(
		log.WithResource(res),
		log.WithProcessor(
			log.NewBatchProcessor(exporter),
		),
	)

	return lp, nil
}
