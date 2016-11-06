package assert

import (
	"errors"
	"fmt"
	"reflect"
)

var ForTesting = errors.New("ForTesting")

/*
Template:

func Is{{ .name }}(t Exam, err error, format string, args ...interface{}) {
	if err == nil {
		return
	}

	hasFmt := HasFmt(format)
	hasMsg := HasArgs(args...)

	if hasFmt && hasMsg {
		t.Logf("{{ .format }}", {{ .arg }}, AsFmt(format, args...))
	} else if hasFmt && !hasMsg {
		t.Logf("{{ .format }}", {{ .arg }}, format)
	} else if !hasFmt && hasMsg {
		t.Logf("{{ .format }}", {{ .arg }}, AsMsg(args...))
	} else if !hasFmt && !hasMsg {
		t.Logf("{{ .format }}", err)
	}
	t.Fail()
}
*/

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

func IsNil(t Exam, a interface{}, args ...string) {
	IsNilf(t, a, "", ToAny(args...)...)
}

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

func IsNotNil(t Exam, a interface{}, args ...interface{}) {
	IsNotNilf(t, a, "", args...)
}
