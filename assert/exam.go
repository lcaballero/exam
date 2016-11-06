package assert

// Exam describes the interface used by assertion functions.
type Exam interface {
	Fail()
	FailNow()
	Log(args ...interface{})
	Logf(format string, args ...interface{})
}
