package event

import (
	"reflect"
)

var Hub = &hub{}

type hub struct {
	eventRepository map[reflect.Type][]Handler
}

func (h *hub) Register(handlers ...Handler) {
	if h.eventRepository == nil {
		h.eventRepository = map[reflect.Type][]Handler{}
	}
	for _, handle := range handlers {
		h.eventRepository[handle.EventType()] = append(h.eventRepository[handle.EventType()], handle)
	}
}

func (h *hub) findHandler(t reflect.Type) []Handler {
	return h.eventRepository[t]
}
