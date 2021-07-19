package so

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"sync"
	"testing"
)

var createPwdRegOnce sync.Once

var pwdReg *regexp.Regexp

func createPwdReg() {
	pwd := os.Getenv("pwd")
	pwdReg = regexp.MustCompile(pwd)
}

func line() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return "[so]LineError runtime.Caller(2) Fail"
	}
	createPwdRegOnce.Do(createPwdReg)
	if pwdReg != nil {
		file = pwdReg.ReplaceAllString(file, "")
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
		t.Errorf("%v : Not True", line())
	}
}

// Check value is false
func False(t *testing.T, check bool) {
	if check {
		t.Errorf("%v : Not False", line())
	}
}

// Check two values is equal
func Equal(t *testing.T, a, b interface{}) {
	if a != b && !reflect.DeepEqual(a, b) {
		t.Errorf("%v : Not Equal: %v == %v", line(), a, b)
	}
}

// Check two values isn't equal
func NotEqual(t *testing.T, a, b interface{}) {
	if a == b || reflect.DeepEqual(a, b) {
		t.Errorf("%v : Equal: %v != %v", line(), a, b)
	}
}

// Check object is empty
// If len([]string) == 0, it's empty
func Empty(t *testing.T, target interface{}) {
	if !IsEmpty(target) {
		t.Errorf("%v : Not Empty: %v", line(), target)
	}
}

// Check object isn't empty
func NotEmpty(t *testing.T, target interface{}) {
	if IsEmpty(target) {
		t.Errorf("%v : Empty: %v", line(), target)
	}
}

// Check object == nil
func Nil(t *testing.T, stack interface{}) {
	if stack != nil {
		t.Errorf("%v : Not Nil: %v", line(), stack)
	}
}

// Check object != nil
func NotNil(t *testing.T, stack interface{}) {
	if stack == nil {
		t.Errorf("%v : Nil", line())
	}
}

// Check error != nil
func Error(t *testing.T, err error) {
	if err == nil {
		t.Errorf("%v : Not Error", line())
	}
}
