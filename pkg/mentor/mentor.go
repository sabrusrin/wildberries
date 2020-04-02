package mentor

// Mentor ...
type Mentor interface {
	Listen() (string, error)
}
type mentor struct {
	mentorHeader string
}

// Returns the plan for theory
func (m *mentor) Listen() (res string, err error) {
	res = m.mentorHeader
	err = nil
	return
}

// NewMentor ...
func NewMentor(s string) Mentor {
	return &mentor{
		mentorHeader: s,
	}
}
