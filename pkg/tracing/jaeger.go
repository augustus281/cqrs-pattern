package tracing

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	cfg "github.com/uber/jaeger-client-go/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"

	"github.com/augustus281/cqrs-pattern/global"
	"github.com/augustus281/cqrs-pattern/pkg/config"
)

// Approach 1: Using jaeger
func NewJaegerTracer(jaegerCfg *config.JaegerConfig) (opentracing.Tracer, io.Closer, error) {
	jeagerConfig := &cfg.Configuration{
		ServiceName: jaegerCfg.ServiceName,

		// "const" sampler is a binary sampling strategy: 0=never sample, 1=always sample.
		Sampler: &cfg.SamplerConfig{
			Type:  "const",
			Param: 1,
		},

		// Log the emitted spans to stdout.
		Reporter: &cfg.ReporterConfig{
			LogSpans:           jaegerCfg.LogSpans,
			LocalAgentHostPort: jaegerCfg.HostPort,
		},
	}

	global.Logger.Info("new jaeger tracer successfully")
	return jeagerConfig.NewTracer(cfg.Logger(jaeger.StdLogger))
}

// Approach 2: Using open-telemtry
func StartTracing() (*trace.TracerProvider, error) {
	headers := map[string]string{
		"content-type": "application/json",
	}
	exporter, err := otlptrace.New(
		context.Background(),
		otlptracehttp.NewClient(
			otlptracehttp.WithEndpoint("localhost:4318"),
			otlptracehttp.WithHeaders(headers),
			otlptracehttp.WithInsecure(),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("creating new exporter: %w", err)
	}

	tracerprovider := trace.NewTracerProvider(
		trace.WithBatcher(
			exporter,
			trace.WithMaxExportBatchSize(trace.DefaultMaxExportBatchSize),
			trace.WithBatchTimeout(trace.DefaultScheduleDelay*time.Millisecond),
			trace.WithMaxExportBatchSize(trace.DefaultMaxExportBatchSize),
		),
		trace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String("product-app"),
			),
		),
	)

	otel.SetTracerProvider(tracerprovider)
	return tracerprovider, nil
}
