package domain

import (
	"context"
	"github.com/zhaoche27/colago/core/event"
)

var EventPublisher = &eventPublisher{}

type eventPublisher struct {
}

func (ep *eventPublisher) Publish(ctx context.Context, e event.Event) error {
	return event.Bus.FireAll(ctx, e)
}
