package assert

// Lengthy provides an interface for any type that provides a way to
// get the length of the collection.
type Lengthy interface {
	Len() int
}

// HasLength checks that an instance which can provide a length is
// greater than 0.
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
