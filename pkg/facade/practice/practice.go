package practice

// Interface to work with practice type
type Practice struct {
	task	string
}

func (p *Practice) Solve(s string) string	{
	p.task = p.task + " " + s + ","
	return "Do the practice tasks given by mentors:" + p.task + "\n"
}

// Constructor for practice type
func NewPractice() Practice	{
	return Practice{}
}
