package assert

type Exam interface {
	Fail()
	Log(args ...interface{})
	Logf(format string, args ...interface{})
}
