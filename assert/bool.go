package assert

import "testing"

func True(t *testing.T, a bool) {
	if a {
		return
	}
	t.Log("expected %v to true, but wasn't", a)
	t.Fail()
}
