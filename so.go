package so

import (
	"errors"
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

// Check object is empty
// This func copy from assert
func IsEmpty(object interface{}) bool {

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
		return IsEmpty(deref)
		// for all other types, compare against the zero value
	default:
		zero := reflect.Zero(objValue.Type())
		return reflect.DeepEqual(object, zero.Interface())
	}
}

// Check value is true
func True(t *testing.T, check bool) {
	if !check {
		t.Errorf("%v Not True", line())
	}
}

// Check value is false
func False(t *testing.T, check bool) {
	if check {
		t.Errorf("%v Not False", line())
	}
}

// Check two values is equal
func Equal(t *testing.T, a, b interface{}) {
	if a != b && !reflect.DeepEqual(a, b) {
		t.Errorf("%v Not Equal: %v == %v", line(), a, b)
	}
}

// Check two values isn't equal
func NotEqual(t *testing.T, a, b interface{}) {
	if a == b || reflect.DeepEqual(a, b) {
		t.Errorf("%v Equal: %v != %v", line(), a, b)
	}
}

// Check object is empty
// If len([]string) == 0, it's empty
func Empty(t *testing.T, target interface{}) {
	if !IsEmpty(target) {
		t.Errorf("%v Not Empty: %v", line(), target)
	}
}

// Check object isn't empty
func NotEmpty(t *testing.T, target interface{}) {
	if IsEmpty(target) {
		t.Errorf("%v Empty: %v", line(), target)
	}
}

// Check object == nil
func Nil(t *testing.T, stack interface{}) {
	if stack != nil {
		t.Errorf("%v Not Nil: %v", line(), stack)
	}
}

// Check object != nil
func NotNil(t *testing.T, stack interface{}) {
	if stack == nil {
		t.Errorf("%v Nil", line())
	}
}

// chekcer is error or func(error) bool
func Error(t *testing.T, err error, checker ...interface{}) {
	if len(checker) == 0 {
		if err == nil {
			t.Errorf("%v Error is Nil", line())
		}
		return
	}

	if checker[0] == nil && err == nil {
		return
	}

	for _, check := range checker {
		switch v := check.(type) {
		case error:
			if !errors.Is(err, v) {
				t.Errorf("%v Error isn't type:%v %v\n", line(), checker, err)
			}
		case func(error) bool:
			if !v(err) {
				t.Errorf("%v Erros checker fail %v\n", line(), err)
			}
		}
	}
}
