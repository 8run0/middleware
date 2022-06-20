package main

import (
	"context"
	"fmt"
	"time"
)

type MyThingStoreTimerFuncer struct {
}

func (mts *MyThingStoreTimerFuncer) calculateTime(startTime time.Time) string {
	diff := time.Now().Sub(startTime)
	return fmt.Sprintf("time taken %s", diff)

}

func (mts *MyThingStoreTimerFuncer) TimeGetThing(f GetThingFunc) GetThingFunc {
	return func(ctx context.Context, id int) *Thing {
		startTime := time.Now()
		defer mts.calculateTimeAndLog(ctx, "getting thing %s", startTime)
		return f.GetThing(ctx, id)
	}
}
func (mts *MyThingStoreTimerFuncer) calculateTimeAndLog(ctx context.Context, str string, startTime time.Time) {
	getTools(ctx).Log.Printf(str, mts.calculateTime(startTime))
}

func (mts *MyThingStoreTimerFuncer) TimeAddThing(f AddThingFunc) AddThingFunc {
	return func(ctx context.Context, thing *Thing) int {
		startTime := time.Now()
		defer mts.calculateTimeAndLog(ctx, "adding thing %s", startTime)
		return f.AddThing(ctx, thing)
	}
}

func (mts *MyThingStoreTimerFuncer) TimeDeleteThing(f DeleteThingFunc) DeleteThingFunc {
	return func(ctx context.Context, id int) bool {
		startTime := time.Now()
		defer mts.calculateTimeAndLog(ctx, "deleting thing %s", startTime)
		return f.DeleteThing(ctx, id)
	}
}
