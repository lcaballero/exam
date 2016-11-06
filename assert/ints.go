package assert

// IsEqIntf asserts that the values a and b are equal along with the
// format interface.
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

// IsEqInt asserts that the values a and b are equal along with the
// format interface.
func IsEqInt(t Exam, a, b int, msg ...string) {
	IsEqIntf(t, a, b, "", ToAny(msg...)...)
}

// GreaterThanf asserts that a is greater than b (along with the string
// format interface).
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

// GreaterThan asserts that a is greater than b (and the format interface).
func GreaterThan(t Exam, a, b int, msg ...interface{}) {
	GreaterThanf(t, a, b, "", msg...)
}

// IsZerof asserts that the value is zero (with message and format args).
func IsZerof(t Exam, a int, format string, msg ...interface{}) {
	IsEqIntf(t, a, 0, format, msg)
}

// IsZero asserts that the value is zero.
func IsZero(t Exam, a int, msg ...interface{}) {
	IsEqIntf(t, a, 0, "", msg...)
}

// IsPosf asserts that the value is greater than 0 (along with the string
// format api).
func IsPosf(t Exam, a int, format string, msg ...interface{}) {
	GreaterThanf(t, a, 0, format, msg...)
}

// IsPos asserts that the value is greater than 0.
func IsPos(t Exam, b int, msg ...string) {
	IsPosf(t, b, "", ToAny(msg...)...)
}
