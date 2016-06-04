package assert

import (
	"reflect"
	"testing"
)

func IsEqFunc(t *testing.T, a, b interface{}) {
	sf1 := reflect.ValueOf(a)
	sf2 := reflect.ValueOf(b)
	eqFuncs := sf1.Pointer() == sf2.Pointer()
	if !eqFuncs {
		t.Fail()
	}
}

func IsNotEqFunc(t *testing.T, a, b interface{}) {
	sf1 := reflect.ValueOf(a)
	sf2 := reflect.ValueOf(b)
	eqFuncs := sf1.Pointer() == sf2.Pointer()
	if eqFuncs {
		t.Fail()
	}
}
