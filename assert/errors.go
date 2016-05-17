package assert

import "testing"

func NilError(t *testing.T, err error) {
	if err == nil {
		return
	}
	t.Log("expected nil error", err)
	t.Fail()
}
