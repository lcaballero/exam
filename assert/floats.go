package assert

// IsEqFloat64f asserts that the two values are equal along with
// the format string signature.
func IsEqFloat64f(t Exam, a, b float64, format string, msg ...interface{}) {
	if a == b {
		return
	}

	hasFmt := HasFmt(format)
	hasMsg := HasArgs(msg...)

	if hasFmt && hasMsg {
		t.Logf("%f == %f (%t), %s", a, b, a == b, AsFmt(format, msg...))
	} else if hasFmt && !hasMsg {
		t.Logf("%f == %f (%t), %s", a, b, a == b, format)
	} else if !hasFmt && hasMsg {
		t.Logf("%f == %f (%t), %s", a, b, a == b, AsMsg(msg...))
	} else if !hasFmt && !hasMsg {
		t.Logf("%f == %f (%t)", a, b, a == b)
	}
	t.Fail()
}

// IsEqFloat64 asserts that the two values are equal along with joining
// the given arguments.
func IsEqFloat64(t Exam, a, b float64, msg ...string) {
	IsEqFloat64f(t, a, b, "", ToAny(msg...)...)
}
