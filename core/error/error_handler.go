package error

import (
	"context"
	"github.com/zhaoche27/colago/common/dto"
)

type ErrorHandler interface {
	Handle(ctx context.Context, command dto.Commander, response *dto.Response, err error)
}

var currentErrorHandler ErrorHandler = &DefaultErrorHandler{}

func SetCurrentErrorHandler(errorHandler ErrorHandler) {
	currentErrorHandler = errorHandler
}

func CurrentErrorHandler() ErrorHandler {
	return currentErrorHandler
}

type DefaultErrorHandler struct {
}

func (deh *DefaultErrorHandler) Handle(ctx context.Context, command dto.Commander, response *dto.Response, err error) {

}
