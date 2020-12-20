// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package errors

import "fmt"

type DefaultError struct {
	Message string
}

func (d DefaultError) Error() string  {
	return d.Message
}

func (d DefaultError) String() string  {
	return d.Message
}

func Bomb(format string, args ...interface{})  {
	panic(DefaultError{Message: fmt.Sprintf(format, args...)})
}

func Dangerous(v interface{})  {
	if v == nil {
		return
	}
	switch t := v.(type) {
	case string:
		if t != "" {
			panic(DefaultError{Message: t})
		}
	case error:
		panic(DefaultError{Message: t.Error()})
	}
}
