package command

import (
	"context"
	"github.com/zhaoche27/colago/common/dto"
	"reflect"
)

type CommandExecutor interface {
	Execute(ctx context.Context, command dto.Commander) (*dto.Response, error)
	CommandType() reflect.Type
}
