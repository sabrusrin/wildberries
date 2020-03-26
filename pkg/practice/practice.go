package practice

// TaskList ...
type TaskList interface {
	AppendTaskList(string) string
}

type taskList struct {
	task           []string
	practiceHeader string
}

// Appends the task list and returns the plan for practice
func (p *taskList) AppendTaskList(s string) string {
	var iter int
	var taskList, buffer string
	if len(s) != 0 {
		p.task = append(p.task, s)
	}
	for iter, buffer = range p.task {
		if iter == 0 {
			taskList += " " + buffer
		} else {
			taskList += ", " +buffer
		}
	}
	return p.practiceHeader + taskList + "\n"
}

// NewPracticeList ...
func NewTaskList(s string) TaskList {
	return &taskList{
		practiceHeader: s,
	}
}
