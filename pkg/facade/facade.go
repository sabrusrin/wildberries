package facade

import (
	"WB/pkg/facade/mentor"
	"WB/pkg/facade/practice"
	"WB/pkg/facade/theory"
)

// Interface to work with trialPeriod type
type Planner interface {
	Plan(string, string) string
}

type trialPeriodPlan struct	{
	theory		theory.Reader
	mentor		mentor.Listener
	practice	practice.Solver
}

func (t *trialPeriodPlan) Plan(s1 string, s2 string) string	{
	res := t.theory.Read(s1) + t.mentor.Listen() + t.practice.Solve(s2)
	return res
}

// Constructor for trialPeriod interface
func NewTrial() Planner {
	return &trialPeriodPlan{
		theory:		theory.NewTheory(),
		mentor:		mentor.NewMentor(),
		practice:	practice.NewPractice(),
	}
}