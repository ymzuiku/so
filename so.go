package so

import (
	"fmt"
	"runtime"
	"testing"
)

func line() string {
	pc, file, line, ok := runtime.Caller(2)
	f := runtime.FuncForPC(pc)
	if !ok {
		return "[so]LineError runtime.Caller(2) Fail"
	}
	return fmt.Sprintf("%s:%d (Method %s)\n", file, line, f.Name())
}

func True(t *testing.T, check bool, msg ...interface{}) {
	if !check {
		t.Errorf("%v %v", line(), msg)
	}
}

func False(t *testing.T, check bool, msg ...interface{}) {
	if check {
		t.Errorf("%v %v", line(), msg)
	}
}

func Nil(t *testing.T, stack interface{}, msg ...interface{}) {
	if stack != nil {
		t.Errorf("%v %v%v", line(), stack, msg)
	}
}

func NotNil(t *testing.T, stack interface{}, msg ...interface{}) {
	if stack == nil {
		t.Errorf("%v %v%v", line(), stack, msg)
	}
}

func Error(t *testing.T, err error, msg ...interface{}) {
	if err == nil {
		t.Errorf("%v %w %v", line(), err, msg)
	}
}
