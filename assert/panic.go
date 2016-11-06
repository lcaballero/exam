package assert

// ShouldPanic asserts that the given function panics.
func ShouldPanic(t Exam, fn func()) {
	defer func() {
		err := recover()
		t.Log(err)
		if err == nil {
			t.Log("Expected panic occurred.")
			t.Fail()
		}
	}()
	fn()
	t.Log("Expecting panic.")
	t.Fail()
}

// ShouldNotPanic asserts that the given function does not panic.
func ShouldNotPanic(t Exam, fn func()) {
	defer func() {
		err := recover()
		if err != nil {
			t.Log("Expect func() not to panic.", err)
			t.Fail()
		}
	}()
	fn()
}
