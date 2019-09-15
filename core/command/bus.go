package command

import (
	"context"
	"fmt"
	"github.com/zhaoche27/colago/common/dto"
	error2 "github.com/zhaoche27/colago/core/error"
	"reflect"
)

var Bus = &bus{}

type bus struct {
}

func (b *bus) Send(ctx context.Context, command dto.Commander) (response *dto.Response) {
	commandType := reflect.TypeOf(command)
	fmt.Println(commandType.String())
	ci, ok := Hub.commandRepository[commandType]
	if !ok {
		error2.CurrentHandler().Handle(ctx, command, response, error2.NewColaErrf("Not found executor, execute `%T` ", command))
		return
	}
	response = ci.invoke(ctx, command)
	return
}
