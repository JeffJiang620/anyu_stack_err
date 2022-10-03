package stackErr

import "runtime"

const DefaultErrorStackSkip = 1

type errorWithStack struct {
	err  error
	file string
	line int
}

func (e errorWithStack) Error() string {
	return e.err.Error()
}

func (e errorWithStack) Unwrap() error {
	return e.err
}

func (e errorWithStack) File() string {
	return e.file
}

func (e errorWithStack) Line() int {
	return e.line
}

func WithStack(err error, skip int) error {
	switch err.(type) {
	case errorWithStack:
		return err
	default:
		_, file, line, ok := runtime.Caller(skip)
		if ok {
			return errorWithStack{
				err:  err,
				file: file,
				line: line,
			}
		}
		return err
	}
}

func StackErr(err error) error {
	return WithStack(err, 2)
}
