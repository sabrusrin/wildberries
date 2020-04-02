package facade

type theory interface {
	Read(string) string
}

type mentor interface {
	Listen() string
}

type practice interface {
	Solve(string) string
}

// Planner Interface to work with trialPeriod type
type Planner interface {
	Plan(string, string) string
}

type trialPeriodPlan struct {
	theory   theory
	mentor   mentor
	practice practice
}

// Plan method returns a plan for trialPeriod
func (t *trialPeriodPlan) Plan(s1 string, s2 string) string {
	res := t.theory.Read(s1) + t.mentor.Listen() + t.practice.Solve(s2)
	return res
}

// NewPlanner for trialPeriod interface
func NewPlanner(
	theory theory,
	mentor mentor,
	practice practice,
	) Planner {
	return &trialPeriodPlan{
		theory:   theory,
		mentor:   mentor,
		practice: practice,
	}
}
