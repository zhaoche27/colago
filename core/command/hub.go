package command

import (
	"context"
	"fmt"
	"github.com/zhaoche27/colago/common/dto"
	error2 "github.com/zhaoche27/colago/core/error"
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

func (h *hub) Send(ctx context.Context, command dto.Commander) (response *dto.Response) {
	commandType := reflect.TypeOf(command)
	fmt.Println(commandType.String())
	ci, ok := h.commandRepository[commandType]
	if !ok {
		error2.CurrentHandler().Handle(ctx, command, response, fmt.Errorf("Not found executor, execute `%T` ", command))
		return
	}
	response = ci.invoke(ctx, command)
	return
}
