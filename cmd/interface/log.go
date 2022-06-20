package main

import "context"

type MyThingStoreLogFuncer struct {
}

func (mts *MyThingStoreLogFuncer) LogGetThing(f GetThingFunc) GetThingFunc {
	return func(ctx context.Context, id int) *Thing {
		getTools(ctx).Log.Printf("getting thing with id: %d", id)
		return f.GetThing(ctx, id)
	}
}
func (mts *MyThingStoreLogFuncer) LogAddThing(f AddThingFunc) AddThingFunc {
	return func(ctx context.Context, thing *Thing) int {
		getTools(ctx).Log.Printf("adding thing with id %v", thing.id)
		return f.AddThing(ctx, thing)
	}
}

func (mts *MyThingStoreLogFuncer) LogDeleteThing(f DeleteThingFunc) DeleteThingFunc {
	return func(ctx context.Context, id int) bool {
		getTools(ctx).Log.Printf("deleting thing with id: %d", id)
		return f.DeleteThing(ctx, id)
	}
}
