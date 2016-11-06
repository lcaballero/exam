package assert

import (
	"strings"
	"testing"
)

// Contains asserts that the given string contains each of the
// substrings.
func Contains(t *testing.T, st string, subs ...string) {
	for _, s := range subs {
		if !strings.Contains(st, s) {
			t.Logf("'%s' string should have 'contained' %s", st, s)
			t.Fail()
		}
	}
}

// Excludes asserts that the given string does not contain any of
// the substrings.
func Excludes(t *testing.T, st string, subs ...string) {
	for _, s := range subs {
		if strings.Contains(st, s) {
			t.Logf("'%s' should NOT have 'contained' %s", st, s)
			t.Fail()
		}
	}
}

// IsEmptyString asserts that the given string is empty.
func IsEmptyString(t Exam, a string) {
	if a == "" {
		return
	}
	t.Logf("'%s' == '' (%t)", a, a == "")
	t.Fail()
}

// IsNotEmptyString asserts that the given string is not empty.
func IsNotEmptyString(t Exam, a string) {
	if a != "" {
		return
	}
	t.Logf("'%s' != '' (%t)", a, a == "")
	t.Fail()
}

// IsEqStrings asserts the two strings are equal.
func IsEqStrings(t Exam, a, b string) {
	if a == b {
		return
	}
	t.Logf("'%s' == '%s' (%t)", a, b, a == b)
	t.Fail()
}
