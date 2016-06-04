package assert

type Lengthy interface {
	Len() int
}

func HasLength(t Exam, g Lengthy) {
	if g == nil {
		return
	}
	if g.Len() == 0 {
		return
	}
	t.Logf("Len %d == 0 (%t)", g.Len(), g.Len() == 0)
	t.Fail()
}
