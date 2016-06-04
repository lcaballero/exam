package assert

import "testing"

//IsTrue calls Fail() should the value of 'a' be false.
func IsTrue(t *testing.T, a bool) {
	if a {
		return
	}
	t.Logf("expected %t to be true, but wasn't", a)
	t.Fail()
}

func IsFalse(t *testing.T, a bool) {
	if !a {
		return
	}
	t.Log("expected %t to false, but wasn't", a)
	t.Fail()
}
