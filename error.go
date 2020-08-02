package retchan

import (
	"fmt"
	"runtime"
)

type Error struct {
	orig interface{}
	Stack []struct{File string; Line int}
}

func (e *Error) Error () string {
	return fmt.Sprintf("%s", e.orig)
}

func NewError (orig interface{}, withStack bool) (err *Error) {
	err = &Error{
		orig: orig,
	}

	if !withStack { return }

	skip := 1

	_,file,line,ok := runtime.Caller(skip)

	for ok {
		err.Stack = append(err.Stack, struct {File string;Line int}{file,line})
		skip++
		_,file,line,ok = runtime.Caller(skip)
	}

	return
}