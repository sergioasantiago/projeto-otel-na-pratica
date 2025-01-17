package telemetry

import (
	"context"
	"log/slog"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/propagation"
)

// InitTelemetry initializes OpenTelemetry SDK with traces, metrics and logging
func InitTelemetry(ctx context.Context, cfg *config.TelemetryConfig, serviceName string) (func(), error) {
	// Initialize logs
	loggerProvider, err := initLoggerProvider(ctx, cfg, serviceName)
	if err != nil {
		return nil, err
	}
	global.SetLoggerProvider(loggerProvider)

	// Initialize tracer
	tracerProvider, err := initTraceProvider(ctx, cfg, serviceName)
	if err != nil {
		return nil, err
	}
	otel.SetTracerProvider(tracerProvider)

	// Initialize metrics
	meterProvider, err := initMetricsProvider(ctx, cfg, serviceName)
	if err != nil {
		return nil, err
	}
	otel.SetMeterProvider(meterProvider)

	// Set global propagator
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	// Return cleanup function
	return func() {
		if err := loggerProvider.Shutdown(ctx); err != nil {
			slog.Error("Failed to shutdown logger provider", "error", err)
		}
		if err := tracerProvider.Shutdown(ctx); err != nil {
			slog.Error("Failed to shutdown tracer provider", "error", err)
		}
		if err := meterProvider.Shutdown(ctx); err != nil {
			slog.Error("Failed to shutdown meter provider", "error", err)
		}
	}, nil
}
