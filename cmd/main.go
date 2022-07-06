package main

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/metric/instrument"
	"go.opentelemetry.io/otel/metric/instrument/syncint64"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"go.opentelemetry.io/otel/trace"
)

const (
	instrumentationName    = "github.com/instrumentron"
	instrumentationVersion = "v0.1.0"
)

type OTELTools struct {
	Ctx     context.Context
	Tracer  trace.Tracer
	Meter   metric.Meter
	Cleanup func()
}

func NewOTELTools(ctx context.Context, meterName string) *OTELTools {
	// Registers a tracer Provider globally.
	cleanup := InstallExportPipeline(ctx)
	tracer := otel.GetTracerProvider().Tracer(
		instrumentationName,
		trace.WithInstrumentationVersion(instrumentationVersion),
		trace.WithSchemaURL(semconv.SchemaURL),
	)
	meter := global.MeterProvider().Meter(meterName)
	return &OTELTools{
		Ctx:     ctx,
		Tracer:  tracer,
		Meter:   meter,
		Cleanup: cleanup,
	}
}

type Calculator struct {
	calculatorImpl
}

type calculatorImpl interface {
	add(x, y int64) int64
	multiply(x, y int64) int64
	sub(x, y int64) int64
	devide(x, y int64) int64
}

var _ calculatorImpl = &calculator{}
var _ calculatorImpl = &calculatorSpanner{}
var _ calculatorImpl = &calculatorRecorder{}

type calculatorRecorder struct {
	*OTELTools
	durRecorder syncint64.Histogram
	next        calculatorImpl
}

// add implements Calculator
func (c *calculatorRecorder) add(x int64, y int64) int64 {
	startTime := time.Now()
	defer c.RecordTime(startTime, "Addition")
	return c.next.add(x, y)
}

// devide implements Calculator
func (c *calculatorRecorder) devide(x int64, y int64) int64 {
	startTime := time.Now()
	defer c.RecordTime(startTime, "Devision")
	return c.next.devide(x, y)
}

// multiply implements Calculator
func (c *calculatorRecorder) multiply(x int64, y int64) int64 {
	startTime := time.Now()
	defer c.RecordTime(startTime, "Multiply")
	return c.next.multiply(x, y)
}

// sub implements Calculator
func (c *calculatorRecorder) sub(x int64, y int64) int64 {
	startTime := time.Now()
	defer c.RecordTime(startTime, "Subtract")
	return c.next.sub(x, y)
}

func (c *calculatorRecorder) RecordTime(startTime time.Time, name string) {
	now := time.Now()
	dur := now.Sub(startTime)
	c.durRecorder.Record(c.Ctx, dur.Milliseconds())
}

type calculatorSpanner struct {
	*OTELTools
	next calculatorImpl
}

// add implements CalculatorImpl
func (c *calculatorSpanner) add(x int64, y int64) int64 {
	ctx, span := c.Tracer.Start(c.Ctx, "Addition")
	c.Ctx = ctx
	defer span.End()
	return c.next.add(x, y)
}

// devide implements CalculatorImpl
func (c *calculatorSpanner) devide(x int64, y int64) int64 {
	ctx, span := c.Tracer.Start(c.Ctx, "Devision")
	c.Ctx = ctx
	defer span.End()
	return c.next.devide(x, y)
}

// multiply implements CalculatorImpl
func (c *calculatorSpanner) multiply(x int64, y int64) int64 {
	ctx, span := c.Tracer.Start(c.Ctx, "Multiply")
	c.Ctx = ctx
	defer span.End()
	return c.next.multiply(x, y)
}

// sub implements CalculatorImpl
func (c *calculatorSpanner) sub(x int64, y int64) int64 {
	ctx, span := c.Tracer.Start(c.Ctx, "Subtract")
	c.Ctx = ctx
	defer span.End()
	return c.next.sub(x, y)
}

type calculator struct {
}

// devide implements CalculatorImpl
func (c *calculator) devide(x, y int64) int64 {
	return x / y
}

// sub implements CalculatorImpl
func (c *calculator) sub(x, y int64) int64 {
	return x - y
}

func (c *calculator) add(x, y int64) int64 {
	return x + y
}

func (c *calculator) multiply(x, y int64) int64 {
	return x * y
}

func Resource() *resource.Resource {
	return resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String("calculator"),
		semconv.ServiceVersionKey.String("0.0.1"),
	)
}

func InstallExportPipeline(ctx context.Context) func() {
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
func NewCalculator(ctx context.Context) (*Calculator, func()) {
	name := "calculator"
	// Registers a tracer Provider globally.
	tools := NewOTELTools(ctx, name)
	durRecorder, _ := tools.Meter.SyncInt64().Histogram(
		name,
		instrument.WithUnit("milliseconds"),
	)
	calc := calculatorSpanner{
		OTELTools: tools,
		next: &calculatorRecorder{
			OTELTools:   tools,
			durRecorder: durRecorder,
			next:        &calculator{},
		},
	}
	return &Calculator{
		calculatorImpl: &calc,
	}, tools.Cleanup
}

func main() {
	calc, cleanup := NewCalculator(context.Background())
	defer cleanup()
	log.Println("the answer is", calc.add(1, 1))
	log.Println("the answer is", calc.sub(10, 1))
	log.Println("the answer is", calc.multiply(1, 10))
	log.Println("the answer is", calc.devide(100, 10))
  }
