package qerrors

import (
	"fmt"
)

type Error struct {
	code    ErrorCode
	message string
	err     error
}

func New(code ErrorCode, message string, err error) *Error {
	return &Error{
		code:    code,
		message: message,
		err:     err,
	}
}

func (e *Error) Error() string {
	if e.err != nil {
		return fmt.Sprintf("[%s] %s: %v", e.code, e.message, e.err)
	}

	return fmt.Sprintf("[%s] %s", e.code, e.message)
}

func (e *Error) Code() ErrorCode {
	return e.code
}

func (e *Error) Msg() string {
	return e.message
}

func (e *Error) Err() error {
	return e.err
}

func (e *Error) Unwrap() error {
	return e.err
}
