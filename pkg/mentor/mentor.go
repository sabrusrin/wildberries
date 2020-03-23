package mentor

type Mentor struct	{
	mentorHeader string
}

// Returns the plan for theory
func (m *Mentor) Listen() string	{
	return m.mentorHeader
}

// NewMentor ...
func NewMentor(s string) Mentor	{
	return Mentor{
		mentorHeader: s,
	}
}
