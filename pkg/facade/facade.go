package facade

type bookList interface {
	AppendBookList(string) string
}

type mentor interface {
	Listen() string
}

type taskList interface {
	AppendTaskList(string) string
}

// TrialPeriodPlanner Interface to work with trialPeriod type
type TrialPeriodPlanner interface {
	Plan(string, string) string
}

type trialPeriodPlan struct {
	theory   bookList
	mentor   mentor
	practice taskList
}

// Plan method returns a plan for trialPeriod
func (t *trialPeriodPlan) Plan(s1 string, s2 string) string {
	res := t.theory.AppendBookList(s1) + t.mentor.Listen() + t.practice.AppendTaskList(s2)
	return res
}

// NewPlanner for trialPeriod interface
func NewPlanner(
	theory bookList,
	mentor mentor,
	practice taskList,
	) TrialPeriodPlanner {
	return &trialPeriodPlan{
		theory:   theory,
		mentor:   mentor,
		practice: practice,
	}
}
