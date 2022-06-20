package main

import (
	"context"
)

type Thing struct {
	id int
}

type GetThingFunc func(ctx context.Context, id int) *Thing

func (g GetThingFunc) GetThing(ctx context.Context, id int) *Thing {
	return g(ctx, id)
}

type AddThingFunc func(ctx context.Context, thing *Thing) int

func (a AddThingFunc) AddThing(ctx context.Context, thing *Thing) int {
	return a(ctx, thing)
}

type DeleteThingFunc func(ctx context.Context, id int) bool

func (d DeleteThingFunc) DeleteThing(ctx context.Context, id int) bool {
	return d(ctx, id)
}

type MyThingStoreGetter interface {
	GetThing(ctx context.Context, id int) *Thing
}
type MyThingStoreAdder interface {
	AddThing(ctx context.Context, thing *Thing) int
}
type MyThingStoreDeleter interface {
	DeleteThing(ctx context.Context, id int) bool
}
type MyThingStore interface {
	MyThingStoreGetter
	MyThingStoreAdder
	MyThingStoreDeleter
}

func main() {
	inmemStore := NewInMemMyThingStore()
	ctx := context.Background()
	inmemStore.tools.InstallExportPipeline(ctx)
	inmemStore.AddThing(ctx, &Thing{
		id: 0,
	})
	inmemStore.GetThing(ctx, 1)
	inmemStore.GetThing(ctx, 1)

}
