package event

import (
	"context"
	"github.com/zhaoche27/colago/common/event"
	error2 "github.com/zhaoche27/colago/core/error"
	"reflect"
)

var Bus = bus{}

type bus struct {
}

func (b *bus) FireAll(ctx context.Context, event event.Event) error {
	hs := Hub.findHandler(reflect.TypeOf(event))
	if len(hs) == 0 {
		return error2.NewColaErrf("Event type `%T`, not find handler", event)
	}
	for _, h := range hs {
		err := h.Execute(ctx, event)
		if err != nil {
			return err
		}
	}
	return nil
}
