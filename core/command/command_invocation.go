package command

import (
	"context"
	"fmt"
	"github.com/zhaoche27/colago/common/dto"
	error2 "github.com/zhaoche27/colago/core/error"
)

type commandInvocation struct {
	commandHub      *commandHub
	commandExecutor CommandExecutor
}

func newCommandInvocation(commandHub *commandHub, commandExecutor CommandExecutor) *commandInvocation {
	return &commandInvocation{commandHub: commandHub, commandExecutor: commandExecutor}
}

func (ci *commandInvocation) invoke(ctx context.Context, command dto.Commander) (response *dto.Response) {
	defer func() {
		r := recover()
		if r != nil {
			switch err := r.(type) {
			case error:
				error2.CurrentErrorHandler().Handle(ctx, command, response, err)
				return
			default:
				error2.CurrentErrorHandler().Handle(ctx, command, response, fmt.Errorf("%v", err))
				return
			}
		}
		err := ci.postIntercept(ctx, command, response)
		if err != nil {
			error2.CurrentErrorHandler().Handle(ctx, command, response, err)
			return
		}
	}()
	err := ci.preIntercept(ctx, command)
	if err != nil {
		error2.CurrentErrorHandler().Handle(ctx, command, response, err)
		return
	}
	response, err = ci.commandExecutor.Execute(ctx, command)
	if err != nil {
		error2.CurrentErrorHandler().Handle(ctx, command, response, err)
		return
	}
	return
}

func (ci *commandInvocation) preIntercept(ctx context.Context, command dto.Commander) error {
	for _, preIntercept := range ci.commandHub.preIntercepts {
		err := preIntercept.PreIntercept(ctx, command)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ci *commandInvocation) postIntercept(ctx context.Context, command dto.Commander, response *dto.Response) error {
	for _, postIntercept := range ci.commandHub.postIntercepts {
		err := postIntercept.PostInterceptor(ctx, command, response)
		if err != nil {
			return err
		}
	}
	return nil
}
