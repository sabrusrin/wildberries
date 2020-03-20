package facade

type Planner interface {
	Plan() string
}

// A structure to represent Facade
type trialPeriod struct	{
	theory		*theory
	coursera	*coursera
	practice	*practice
}

// Method for a TrialPeriod which returns a plan for trial period
func (t *trialPeriod) Plan() string	{
	res := t.theory.read() + t.coursera.watch() + t.practice.solve()
	return res
}

// Constructor
func NewTrial() Planner {
	return &trialPeriod{
		theory:		&theory{},
		coursera:	&coursera{},
		practice:	&practice{},
	}
}