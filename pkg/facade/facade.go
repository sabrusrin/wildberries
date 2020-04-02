package facade

type bookList interface {
	AppendBookList(string) (string, error)
}

type mentor interface {
	Listen() (string, error)
}

type taskList interface {
	AppendTaskList(string) (string, error)
}

// TrialPeriodPlanner Interface to work with trialPeriod type
type TrialPeriodPlanner interface {
	Plan(string, string) (string, error)
}

type trialPeriodPlan struct {
	theory   bookList
	mentor   mentor
	practice taskList
}

// Plan method returns a plan for trialPeriod
func (t *trialPeriodPlan) Plan(s1 string, s2 string) (res string, err error) {
	var tmp string
	if tmp, err = t.theory.AppendBookList(s1); err == nil {
		res += tmp
		if tmp, err = t.mentor.Listen(); err == nil {
			res += tmp
			if tmp, err = t.practice.AppendTaskList(s2); err == nil {
				res += tmp
			}
		}
	}
	return
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
