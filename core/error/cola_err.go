package error

import "fmt"

type ColaErr struct {
	msg string
}

func NewColaErr(msg string) *ColaErr {
	return &ColaErr{msg: msg}
}

func NewColaErrf(format string, args ...interface{}) *ColaErr {
	return NewColaErr(fmt.Sprintf(format, args...))
}

func (c *ColaErr) Error() string {
	return c.msg
}
