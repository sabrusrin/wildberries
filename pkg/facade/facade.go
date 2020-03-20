package facade

import (
	"WB/pkg/facade/mentor"
	"WB/pkg/facade/practice"
	"WB/pkg/facade/theory"
)

// Interface to work with trialPeriod type
type Planner interface {
	Plan() string
}

type trialPeriod struct	{
	theory		theory.Reader
	mentor		mentor.Listener
	practice	practice.Solver
}

func (t *trialPeriod) Plan() string	{
	res := t.theory.Read() + t.mentor.Listen() + t.practice.Solve()
	return res
}

// Constructor for trialPeriod
func NewTrial() Planner {
	return &trialPeriod{
		theory:		theory.NewTheory(),
		mentor:		mentor.NewMentor(),
		practice:	practice.NewPractice(),
	}
}