package main

import (
	"context"
	"io"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
)

var (
	tracer = otel.GetTracerProvider().Tracer(
		"instrumentationName",
		trace.WithInstrumentationVersion("instrumentationVersion"),
		trace.WithSchemaURL(semconv.SchemaURL),
	)
)

type Tools struct {
	Log  *log.Logger
	Span *SpanTools
}

type SpanTools struct {
	tracer trace.Tracer
}

type LogWriter struct {
	log *log.Logger
}

// Write implements io.Writer
func (l LogWriter) Write(p []byte) (n int, err error) {
	l.log.Print(string(p))
	return 0, nil
}

var _ io.Writer = LogWriter{}

func getTools(ctx context.Context) *Tools {
	tools := ctx.Value("tools").(*Tools)
	return tools
}

type MyThingStoreToolBuilder struct {
}

func (mts *MyThingStoreToolBuilder) AddToolsToContext(ctx context.Context) context.Context {

	tracer := otel.GetTracerProvider().Tracer(
		"instrumentationName",
		trace.WithInstrumentationVersion("instrumentationVersion"),
		trace.WithSchemaURL(semconv.SchemaURL),
	)
	logger := log.Default()
	tools := &Tools{
		Log: logger,
		Span: &SpanTools{
			tracer: tracer,
		},
	}
	ctx = context.WithValue(ctx, "tools", tools)
	mts.InstallExportPipeline(ctx)
	return ctx
}
func Resource() *resource.Resource {
	return resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String("stdout-example"),
		semconv.ServiceVersionKey.String("0.0.1"),
	)
}

func (mts *MyThingStoreToolBuilder) InstallExportPipeline(ctx context.Context) func() {
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		log.Fatalf("creating stdout exporter: %v", err)
	}

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(Resource()),
	)
	otel.SetTracerProvider(tracerProvider)

	return func() {
		if err := tracerProvider.Shutdown(ctx); err != nil {
			log.Fatalf("stopping tracer provider: %v", err)
		}
	}
}
