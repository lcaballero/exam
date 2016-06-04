package assert

func IsEqRunes(ex Exam, a, b rune) {
	if a == b {
		return
	}
	ex.Fail()
}
