package assert

import (
	"errors"
	"fmt"
	"reflect"
)

// ErrForTesting is for use in tests when a specific error can be
// compared to or against.
var ErrForTesting = errors.New("error for testing")

func isNil(a interface{}) (isReallyNil bool) {
	// Eliminates a string with characters '<nil>'
	_, ok := a.(string)
	if ok {
		return false
	}
	s := fmt.Sprintf("%v", a)
	if s == "<nil>" {
		return true
	}
	v := reflect.ValueOf(a)
	if v.IsNil() {
		return true
	}
	return false
}

// IsNilf checks to determine if a is nil as well as concatenating the
// format args.
func IsNilf(t Exam, a interface{}, format string, args ...interface{}) {
	if isNil(a) {
		return
	}

	hasFmt := HasFmt(format)
	hasMsg := HasArgs(args...)

	if hasFmt && hasMsg {
		t.Logf("%v == nil (%t) %s", a, a == nil, AsFmt(format, args...))
	} else if hasFmt && !hasMsg {
		t.Logf("%v == nil (%t) %s", a, a == nil, format)
	} else if !hasFmt && hasMsg {
		t.Logf("%v == nil (%t) %s", a, a == nil, AsMsg(args...))
	} else if !hasFmt && !hasMsg {
		t.Logf("%v == nil (%t)", a, a == nil)
	}
	t.Fail()
}

// IsNil passes if the value is nil (and passes otherwise).
func IsNil(t Exam, a interface{}, args ...string) {
	IsNilf(t, a, "", ToAny(args...)...)
}

// IsNotNilf passes if a is not nil along with the format interface.
func IsNotNilf(t Exam, a interface{}, format string, args ...interface{}) {
	if !isNil(a) {
		return
	}

	hasFmt := HasFmt(format)
	hasMsg := HasArgs(args...)

	if hasFmt && hasMsg {
		t.Logf("%v != nil (%t) %s", a, a != nil, AsFmt(format, args...))
	} else if hasFmt && !hasMsg {
		t.Logf("%v != nil (%t) %s", a, a != nil, format)
	} else if !hasFmt && hasMsg {
		t.Logf("%v != nil (%t) %s", a, a != nil, AsMsg(args...))
	} else if !hasFmt && !hasMsg {
		t.Logf("%v != nil (%t)", a, a != nil)
	}
	t.Fail()
}

// IsNotNil tests that the value is not nil and passes if that
// condition is true.
func IsNotNil(t Exam, a interface{}, args ...interface{}) {
	IsNotNilf(t, a, "", args...)
}
