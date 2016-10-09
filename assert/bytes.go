package assert

import "bytes"

func IsEqBytes(ex Exam, a, b []byte) {
	if bytes.Equal(a, b) {
		return
	}
	ex.Fail()
}

func IsEqByte(ex Exam, a, b byte) {
	if a == b {
		return
	}
	ex.Fail()
}
