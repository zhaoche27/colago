package command

import (
	"reflect"
)

var Hub = &hub{}

type hub struct {
	preIntercepts     []PreInterceptor
	postIntercepts    []PostInterceptor
	commandRepository map[reflect.Type]*invocation
}

func (h *hub) AddPreIntercepts(interceptors ...PreInterceptor) {
	h.preIntercepts = append(h.preIntercepts, interceptors...)
}

func (h *hub) AddPostIntercepts(interceptors ...PostInterceptor) {
	h.postIntercepts = append(h.postIntercepts, interceptors...)
}

func (h *hub) PutCommandExecutor(executor Executor) {
	if h.commandRepository == nil {
		h.commandRepository = make(map[reflect.Type]*invocation)
	}
	ci := newInvocation(h, executor)
	h.commandRepository[executor.CommandType()] = ci
}
