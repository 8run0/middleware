package calculator

import (
	"context"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

type OTELTools struct {
	Ctx     context.Context
	Tracer  trace.Tracer
	Meter   metric.Meter
	Cleanup func()
}

type calculators interface {
	add(x, y int64) int64
	multiply(x, y int64) int64
	sub(x, y int64) int64
	devide(x, y int64) int64
}
