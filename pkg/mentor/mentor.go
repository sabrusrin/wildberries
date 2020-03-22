package mentor

type Mentor struct	{
	mentorHeader string
}

// Returns the plan for theory
func (m *Mentor) Listen() string	{
	return m.mentorHeader
}

// Constructor for mentor type
func NewMentor(s string) Mentor	{
	return Mentor{
		mentorHeader: s,
	}
}
