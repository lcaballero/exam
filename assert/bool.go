package assert

import "testing"

//IsTrue calls Fail() should the value of 'a' be false.
func IsTrue(t *testing.T, a bool) {
	if a {
		return
	}
	t.Logf("expected %t == true (%t), but doesn't", a, a == true)
	t.Fail()
}

func IsFalse(t *testing.T, a bool) {
	if !a {
		return
	}
	t.Logf("expected %t == false (%t), but doesn't", a, a == false)
	t.Fail()
}
