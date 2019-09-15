package error

import (
	"context"
	"github.com/zhaoche27/colago/common/dto"
)

type Handler interface {
	Handle(ctx context.Context, command dto.Commander, response *dto.Response, err error)
}

var currentHandler Handler = &DefaultHandler{}

func SetCurrentHandler(errorHandler Handler) {
	currentHandler = errorHandler
}

func CurrentHandler() Handler {
	return currentHandler
}

type DefaultHandler struct {
}

func (deh *DefaultHandler) Handle(ctx context.Context, command dto.Commander, response *dto.Response, err error) {
	panic("Please SetCurrentHandler")
}
