package main

import (
	"context"
	"time"
)

var _ MyThingStore = &InMemoryMyThingStore{}

type InMemoryMyThingStore struct {
	tools MyThingStoreToolBuilder
	span  MyThingStoreSpanFuncer
	log   MyThingStoreLogFuncer
	timer MyThingStoreTimerFuncer
}

func NewInMemMyThingStore() *InMemoryMyThingStore {
	return &InMemoryMyThingStore{
		tools: MyThingStoreToolBuilder{},
		span:  MyThingStoreSpanFuncer{},
		log:   MyThingStoreLogFuncer{},
		timer: MyThingStoreTimerFuncer{},
	}
}

// AddThing implements MyThingStore
func (i *InMemoryMyThingStore) AddThing(ctx context.Context, thing *Thing) int {
	ctx = i.tools.AddToolsToContext(ctx)
	return i.log.LogAddThing(
		i.span.SpanAddThing(
			i.timer.TimeAddThing(i.addThing)))(ctx, thing)
}

func (i *InMemoryMyThingStore) addThing(ctx context.Context, thing *Thing) int {
	getTools(ctx).Log.Printf("adding thing to in mem store %v", thing)
	time.Sleep(time.Second)
	return thing.id
}

// DeleteThing implements MyThingStore
func (i *InMemoryMyThingStore) DeleteThing(ctx context.Context, id int) bool {
	ctx = i.tools.AddToolsToContext(ctx)
	return i.log.LogDeleteThing(
		i.span.SpanDeleteThing(
			i.timer.TimeDeleteThing(i.deleteThing)))(ctx, id)
}

func (i *InMemoryMyThingStore) deleteThing(ctx context.Context, id int) bool {
	getTools(ctx).Log.Printf("deleting thing from in mem store with id %d", id)
	time.Sleep(time.Second)
	return true
}

// GetThing implements MyThingStore
func (i *InMemoryMyThingStore) GetThing(ctx context.Context, id int) *Thing {
	ctx = i.tools.AddToolsToContext(ctx)
	return i.log.LogGetThing(
		i.span.SpanGetThing(
			i.timer.TimeGetThing(i.getThing)))(ctx, id)
}
func (i *InMemoryMyThingStore) getThing(ctx context.Context, id int) *Thing {
	getTools(ctx).Log.Printf("getting thing from in mem store with id %d", id)
	return &Thing{
		id: id,
	}
}
