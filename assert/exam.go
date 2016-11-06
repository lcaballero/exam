package assert

type Exam interface {
	Fail()
	FailNow()
	Log(args ...interface{})
	Logf(format string, args ...interface{})
}
