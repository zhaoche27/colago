package command

import (
	"context"
	"fmt"
	"github.com/zhaoche27/colago/common/dto"
	error2 "github.com/zhaoche27/colago/core/error"
	"reflect"
)

var CommandHub = &commandHub{}

type commandHub struct {
	preIntercepts     []PreInterceptor
	postIntercepts    []PostInterceptor
	commandRepository map[reflect.Type]*commandInvocation
}

func (cb *commandHub) AddPreIntercepts(interceptors ...PreInterceptor) {
	cb.preIntercepts = append(cb.preIntercepts, interceptors...)
}

func (cb *commandHub) AddPostIntercepts(interceptors ...PostInterceptor) {
	cb.postIntercepts = append(cb.postIntercepts, interceptors...)
}

func (cb *commandHub) PutCommandExecutor(executor CommandExecutor) {
	if cb.commandRepository == nil {
		cb.commandRepository = make(map[reflect.Type]*commandInvocation)
	}
	ci := newCommandInvocation(cb, executor)
	cb.commandRepository[executor.CommandType()] = ci
}

func (cb *commandHub) Send(ctx context.Context, command dto.Commander) (response *dto.Response) {
	commandType := reflect.TypeOf(command)
	fmt.Println(commandType.String())
	ci, ok := cb.commandRepository[commandType]
	if !ok {
		error2.CurrentErrorHandler().Handle(ctx, command, response, fmt.Errorf("Not found executor, execute `%T` ", command))
		return
	}
	response = ci.invoke(ctx, command)
	return
}
