package practice

type Practice struct	{
	task	string
	practiceHeader string
}

// Appends the task list and returns the plan for practice
func (p *Practice) Solve(s string) string	{
	if len(s) != 0	{
		if len(p.task) == 0	{
			p.task = p.task + " " + s
		} else	{
			p.task = p.task + ", " + s
		}
	}
	return p.practiceHeader + p.task + "\n"
}

// Constructor for practice type
func NewPractice(s string) Practice	{
	return Practice{
		practiceHeader: s,
	}
}
