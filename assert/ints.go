package assert

func IsEqIntf(t Exam, a, b int, format string, msg ...interface{}) {
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

func IsEqInt(t Exam, a, b int, msg ...string) {
	IsEqIntf(t, a, b, "", ToAny(msg...)...)
}

func GreaterThanf(t Exam, a, b int, format string, msg ...interface{}) {
	if a > b {
		return
	}

	hasFmt := HasFmt(format)
	hasMsg := HasArgs(msg...)

	if hasFmt && hasMsg {
		t.Logf("%d > %d (%t), %s", a, b, a > b, AsFmt(format, msg...))
	} else if hasFmt && !hasMsg {
		t.Logf("%d > %d (%t), %s", a, b, a > b, format)
	} else if !hasFmt && hasMsg {
		t.Logf("%d > %d (%t), %s", a, b, a > b, AsMsg(msg...))
	} else if !hasFmt && !hasMsg {
		t.Logf("%d > %d (%t)", a, b, a > b)
	}
	t.Fail()
}

func GreaterThan(t Exam, a, b int, msg ...interface{}) {
	GreaterThanf(t, a, b, "", msg...)
}

func IsZerof(t Exam, a int, format string, msg ...interface{}) {
	IsEqIntf(t, a, 0, format, msg)
}
func IsZero(t Exam, a int, msg ...interface{}) {
	IsEqIntf(t, a, 0, "", msg...)
}

func IsPosf(t Exam, a int, format string, msg ...interface{}) {
	GreaterThanf(t, a, 0, format, msg...)
}

func IsPos(t Exam, b int, msg ...string) {
	IsPosf(t, b, "", ToAny(msg...)...)
}
