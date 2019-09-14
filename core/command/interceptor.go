package command

import (
	"context"
	"github.com/zhaoche27/colago/common/dto"
)

type PreInterceptor interface {
	PreIntercept(ctx context.Context, command dto.Commander) error
}

type PostInterceptor interface {
	PostInterceptor(ctx context.Context, command dto.Commander, response *dto.Response) error
}
