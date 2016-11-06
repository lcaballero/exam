package assert

import "bytes"

// IsEqBytes compares to slices of bytes.
func IsEqBytes(ex Exam, a, b []byte) {
	if bytes.Equal(a, b) {
		return
	}
	ex.Fail()
}

// IsEqByte compares the two bytes for equality.
func IsEqByte(ex Exam, a, b byte) {
	if a == b {
		return
	}
	ex.Fail()
}
