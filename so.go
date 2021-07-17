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

func True(t *testing.T, check bool, msg interface{}) {
	if !check {
		t.Errorf("%v%v", line(), msg)
	}
}

func False(t *testing.T, check bool, msg interface{}) {
	if check {
		t.Errorf("%v%v", line(), msg)
	}
}

func Nil(t *testing.T, stack interface{}) {
	if stack != nil {
		t.Errorf("%v%v", line(), stack)
	}
}

func Error(t *testing.T, err error) {
	if err == nil {
		t.Errorf("%v%w", line(), err)
	}
}
