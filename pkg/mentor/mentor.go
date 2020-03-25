package mentor

// Mentor ...
type Mentor interface {
	Listen() string
}
type mentor struct {
	mentorHeader string
}

// Returns the plan for theory
func (m *mentor) Listen() string {
	return m.mentorHeader
}

// NewMentor ...
func NewMentor(s string) Mentor {
	return &mentor{
		mentorHeader: s,
	}
}
