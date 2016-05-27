package assert
import "bytes"

func IsEqBytes(ex Exam, a, b []byte) {
	if bytes.Equal(a, b) {
		return
	}
	ex.Fail()
}
