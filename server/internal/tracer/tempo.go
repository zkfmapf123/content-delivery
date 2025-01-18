package tracer

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type TracerParmas struct {
	servername string
	version string
	port string
	brokers string
	tempo_url string
}

func InitTracer(servername, version, port, brokers, tempo_url string) TracerParmas {

	return TracerParmas{
		servername: servername,
		version: version,
		port: port,
		brokers: brokers,
		tempo_url: tempo_url,
	}
}

func (t TracerParmas) SetTracer(context.Context) func(context.Context) error {
	ctx := context.Background()

	exporter ,err := otlptracegrpc.New(ctx, 
		otlptracegrpc.WithEndpoint(t.tempo_url),
		otlptracegrpc.WithInsecure(),
	)

	if err != nil {
		log.Fatalf("Failed to create exporter: %v", err)
	}

	resources, err := resource.New(
		ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String(t.servername),
			semconv.ServiceVersionKey.String(t.version),
		),
	)

	if err != nil {
		log.Fatalf("Failed to create resources: %v", err)
	}

	provider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resources),
	)

	otel.SetTracerProvider(provider)

	return func(ctx context.Context) error {
		ctx, cancel := context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		return provider.Shutdown(ctx)
	}
}

