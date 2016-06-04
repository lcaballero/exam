package assert

func ShouldPanic(t Exam, fn func()) {
	t.Log("Expecting panic error.")
	defer func() {
		t.Log("Expecting panic error.")
		err := recover()
		t.Log(err)
		if err == nil {
			t.Log("Expected panic occured.")
			t.Fail()
		}
	}()
	//	fn()
	t.Log("Expected panic.")
	//	t.Fail()
}

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
