package errors

import "runtime"

const DefaultErrorStackSkip = 1

type ErrorWithStack struct {
	err  error
	file string
	line int
}

func (e ErrorWithStack) Error() string {
	return e.err.Error()
}

func (e ErrorWithStack) Unwrap() error {
	return e.err
}

func (e ErrorWithStack) File() string {
	return e.file
}

func (e ErrorWithStack) Line() int {
	return e.line
}

func WithStack(err error, skip int) error {
	switch err.(type) {
	case ErrorWithStack:
		return err
	default:
		_, file, line, ok := runtime.Caller(skip)
		if ok {
			return ErrorWithStack{
				err:  err,
				file: file,
				line: line,
			}
		}
		return err
	}
}
