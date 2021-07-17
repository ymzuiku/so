package so

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

func line() string {
	_, file, line, ok := runtime.Caller(2)
	// f := runtime.FuncForPC(pc)
	if !ok {
		return "[so]LineError runtime.Caller(2) Fail"
	}
	return fmt.Sprintf("%s:%d\n", file, line)
}

// this func copy from assert
func isEmpty(object interface{}) bool {

	// get nil case out of the way
	if object == nil {
		return true
	}

	objValue := reflect.ValueOf(object)

	switch objValue.Kind() {
	// collection types are empty when they have no element
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		return objValue.Len() == 0
		// pointers are empty if nil or if the value they point to is empty
	case reflect.Ptr:
		if objValue.IsNil() {
			return true
		}
		deref := objValue.Elem().Interface()
		return isEmpty(deref)
		// for all other types, compare against the zero value
	default:
		zero := reflect.Zero(objValue.Type())
		return reflect.DeepEqual(object, zero.Interface())
	}
}

func True(t *testing.T, check bool) {
	if !check {
		t.Errorf("%vNot True:\n", line())
	}
}

func False(t *testing.T, check bool) {
	if check {
		t.Errorf("%vNot False:\n", line())
	}
}

func Equal(t *testing.T, a, b interface{}) {
	if a != b && !reflect.DeepEqual(a, b) {
		t.Errorf("%vNot Equal:\n%v\n%v", line(), a, b)
	}
}

func NotEqual(t *testing.T, a, b interface{}) {
	if a == b || reflect.DeepEqual(a, b) {
		t.Errorf("%vEqual:\n%v\n%v", line(), a, b)
	}
}

func Empty(t *testing.T, target interface{}) {
	if !isEmpty(target) {
		t.Errorf("%vNot Empty:\n%v", line(), target)
	}
}

func NotEmpty(t *testing.T, target interface{}) {
	if isEmpty(target) {
		t.Errorf("%vEmpty:\n%v", line(), target)
	}
}

func Nil(t *testing.T, stack interface{}) {
	if stack != nil {
		t.Errorf("%vNot Nil:\n%v", line(), stack)
	}
}

func NotNil(t *testing.T, stack interface{}) {
	if stack == nil {
		t.Errorf("%vNil:\n%v", line(), stack)
	}
}

func Error(t *testing.T, err error) {
	if err == nil {
		t.Errorf("%vNot Error:\n%w", line(), err)
	}
}
