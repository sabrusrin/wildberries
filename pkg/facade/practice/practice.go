package practice

// Interface to work with practice type
type Solver interface {
	Solve(string) string
}

type practice struct {
	task	string
}

func (p *practice) Solve(s string) string	{
	p.task = p.task + " " + s + ","
	return "Do the practice tasks given by mentors:" + p.task + "\n"
}

// Constructor for practice type
func NewPractice() Solver	{
	return &practice{}
}
