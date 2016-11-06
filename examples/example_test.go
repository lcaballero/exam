package examples_test

import (
	. "github.com/lcaballero/exam/assert"
	"testing"
)

type mystruct struct{}

func Test_Example_1(t *testing.T) {
	t.Logf("Here an example of the assert interface.")

	n := 0
	t.Log("example of IsZero")
	IsZero(t, n)

	n++
	t.Log("example of IsEqInt")
	IsEqInt(t, n, 1)

	n++
	t.Log("example of greater than")
	GreaterThan(t, n, 1)

	n++
	t.Logf("example of IsTrue (%d == 2)", n)
	IsTrue(t, n == 3)

	n++
	t.Log("example of IsFalse")
	IsFalse(t, n == 3)

	a := []byte{'a', 'b', 'c'}
	b := []byte("abc")
	t.Log("example of IsEqBytes")
	IsEqBytes(t, a, b)

	t.Log("example of IsEqByte")
	IsEqByte(t, a[1], b[1])

	my := &mystruct{}
	IsNotNil(t, my)

	var you *mystruct
	IsNil(t, you)

	willPanic := func() {
		panic("panicking... now...")
	}
	ShouldPanic(t, willPanic)

	dontPanic := func() {
		// not pannicking now
	}
	ShouldNotPanic(t, dontPanic)
}
