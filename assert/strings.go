package assert

import (
	"strings"
	"testing"
)

func Contains(t *testing.T, st string, subs ...string) {
	for _, s := range subs {
		if !strings.Contains(st, s) {
			t.Logf("'%s' string should have 'contained' %s", st, s)
			t.Fail()
		}
	}
}

func Excludes(t *testing.T, st string, subs ...string) {
	for _, s := range subs {
		if strings.Contains(st, s) {
			t.Logf("'%s' should NOT have 'contained' %s", st, s)
			t.Fail()
		}
	}
}

func IsEmptyString(t Exam, a string) {
	if a == "" {
		return
	}
	t.Logf("'%s' == '' (%t)", a, a == "")
	t.Fail()
}

func IsEqStrings(t Exam, a, b string) {
	if a == b {
		return
	}
	t.Logf("'%s' == '%s' (%t)", a, b, a == b)
	t.Fail()
}