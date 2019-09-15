package event

import (
	"context"
	"reflect"
)

type Handler interface {
	Execute(ctx context.Context, event Event) error
	EventType() reflect.Type
}
