package so

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

func line() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return "[so]LineError runtime.Caller(2) Fail"
	}
	return fmt.Sprintf("\n%s:%d", file, line)
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
		t.Errorf("%v : Not True", line())
	}
}

func False(t *testing.T, check bool) {
	if check {
		t.Errorf("%v : Not False", line())
	}
}

func Equal(t *testing.T, a, b interface{}) {
	if a != b && !reflect.DeepEqual(a, b) {
		t.Errorf("%v : Not Equal: %v == %v", line(), a, b)
	}
}

func NotEqual(t *testing.T, a, b interface{}) {
	if a == b || reflect.DeepEqual(a, b) {
		t.Errorf("%v : Equal: %v != %v", line(), a, b)
	}
}

func Empty(t *testing.T, target interface{}) {
	if !isEmpty(target) {
		t.Errorf("%v : Not Empty: %v", line(), target)
	}
}

func NotEmpty(t *testing.T, target interface{}) {
	if isEmpty(target) {
		t.Errorf("%v : Empty: %v", line(), target)
	}
}

func Nil(t *testing.T, stack interface{}) {
	if stack != nil {
		t.Errorf("%v : Not Nil", line())
	}
}

func NotNil(t *testing.T, stack interface{}) {
	if stack == nil {
		t.Errorf("%v : Nil: %v", line(), stack)
	}
}

func Error(t *testing.T, err error) {
	if err == nil {
		t.Errorf("%v : Not Error", line())
	}
}
