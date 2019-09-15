package event

import (
	"context"
	"github.com/zhaoche27/colago/common/event"
	"reflect"
)

type Handler interface {
	Execute(ctx context.Context, event event.Event) error
	EventType() reflect.Type
}
