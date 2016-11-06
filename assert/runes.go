package assert

// IsEqRunes compares the two runes and asserts that they equal.
func IsEqRunes(ex Exam, a, b rune) {
	if a == b {
		return
	}
	ex.Fail()
}
