package domain

import (
	"context"
	event2 "github.com/zhaoche27/colago/core/event"
)

var EventPublisher = &eventPublisher{}

type eventPublisher struct {
}

func (ep *eventPublisher) Publish(ctx context.Context, event event2.Event) error {
	return event2.Bus.FireAll(ctx, event)
}
