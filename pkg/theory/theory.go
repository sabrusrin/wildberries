package theory

// Theory ...
type Theory interface {
	Read(string) string
}

type theory struct {
	book         string
	theoryHeader string
}

// Appends the book list and returns the plan for theory
func (t *Theory) Read(s string) string {
	if len(s) != 0 {
		if len(t.book) == 0 {
			t.book = t.book + " " + s
		} else {
			t.book = t.book + ", " + s
		}
	}
	return t.theoryHeader + t.book + "\n"
}

// NewTheory ...
func NewTheory(s string) Theory {
	return &theory{
		theoryHeader: s,
	}
}
