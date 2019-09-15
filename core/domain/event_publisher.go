package domain

import (
	"context"
	"github.com/zhaoche27/colago/common/event"
	event2 "github.com/zhaoche27/colago/core/event"
)

var EventPublisher = &eventPublisher{}

type eventPublisher struct {
}

func (ep *eventPublisher) Publish(ctx context.Context, event event.Event) error {
	return event2.Bus.FireAll(ctx, event)
}
