package practice

// Interface to work with practice type
type Solver interface {
	Solve() string
}

type practice struct {
}

func (p *practice) Solve() string	{
	return "Do the practice tasks given by mentors\n"
}

// Constructor for practice type
func NewPractice() Solver	{
	return &practice{}
}
