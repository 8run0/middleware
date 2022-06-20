package main

import (
	"context"
)

type MyThingStoreSpanFuncer struct {
}

func (mts *MyThingStoreSpanFuncer) SpanGetThing(f GetThingFunc) GetThingFunc {
	return func(ctx context.Context, id int) *Thing {
		ctx, span := tracer.Start(ctx, "getting span started")
		defer span.End()
		return f(ctx, id)
	}
}

func (mts *MyThingStoreSpanFuncer) SpanAddThing(f AddThingFunc) AddThingFunc {
	return func(ctx context.Context, thing *Thing) int {
		ctx, span := tracer.Start(ctx, "adding span started")
		defer span.End()
		return f(ctx, thing)
	}
}

func (mts *MyThingStoreSpanFuncer) SpanDeleteThing(f DeleteThingFunc) DeleteThingFunc {
	return func(ctx context.Context, id int) bool {
		ctx, span := getTools(ctx).Span.tracer.Start(ctx, "deleting span started")
		defer span.End()
		return f(ctx, id)
	}
}
