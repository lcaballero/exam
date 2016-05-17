package assert

func EqIntf(t Exam, a, b int, format string, msg ...interface{}) {
	if a == b {
		return
	}

	hasFmt := HasFmt(format)
	hasMsg := HasArgs(msg...)

	if hasFmt && hasMsg {
		t.Logf("%d == %d (%v), %s", a, b, a == b, AsFmt(format, msg...))
	} else if hasFmt && !hasMsg {
		t.Logf("%d == %d (%v), %s", a, b, a == b, format)
	} else if !hasFmt && hasMsg {
		t.Logf("%d == %d (%v), %s", a, b, a == b, AsMsg(msg...))
	} else if !hasFmt && !hasMsg {
		t.Logf("%d == %d (%v)", a, b, a == b)
	}
	t.Fail()
}

func EqInt(t Exam, a, b int, msg ...string) {
	if a == b {
		return
	}
	EqIntf(t, a, b, "", ToAny(msg...)...)
}

func GreaterThanf(t Exam, a, b int, format string, msg ...interface{}) {
	if a > b {
		return
	}

	hasFmt := HasFmt(format)
	hasMsg := HasArgs(msg...)

	if hasFmt && hasMsg {
		t.Logf("%d > %d (%v), %s", a, b, a > b, AsFmt(format, msg...))
	} else if hasFmt && !hasMsg {
		t.Logf("%d > %d (%v), %s", a, b, a > b, format)
	} else if !hasFmt && hasMsg {
		t.Logf("%d > %d (%v), %s", a, b, a > b, AsMsg(msg...))
	} else if !hasFmt && !hasMsg {
		t.Logf("%d > %d (%v)", a, b, a > b)
	}
	t.Fail()
}


func Posf(t Exam, a int, format string, msg ...interface{}) {
	GreaterThanf(t, a, 0, format, msg...)
}

func Pos(t Exam, b int, msg ...string) {
	Posf(t, b, "", ToAny(msg...)...)
}

func ToAny(args ...string) []interface{} {
	if args == nil || len(args) < 1 {
		return make([]interface{}, 0)
	}
	any := make([]interface{}, len(args))
	for i,arg := range args {
		any[i] = arg
	}
	return any
}
