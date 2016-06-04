package assert

func IsEqFloat64f(t Exam, a, b float64, format string, msg ...interface{}) {
	if a == b {
		return
	}

	hasFmt := HasFmt(format)
	hasMsg := HasArgs(msg...)

	if hasFmt && hasMsg {
		t.Logf("%d == %d (%t), %s", a, b, a == b, AsFmt(format, msg...))
	} else if hasFmt && !hasMsg {
		t.Logf("%d == %d (%t), %s", a, b, a == b, format)
	} else if !hasFmt && hasMsg {
		t.Logf("%d == %d (%t), %s", a, b, a == b, AsMsg(msg...))
	} else if !hasFmt && !hasMsg {
		t.Logf("%d == %d (%t)", a, b, a == b)
	}
	t.Fail()
}

func IsEqFloat64(t Exam, a, b float64, msg ...string) {
	IsEqFloat64f(t, a, b, "", ToAny(msg...)...)
}
