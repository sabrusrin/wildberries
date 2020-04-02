package facade_test

import (
	"testing"

	"github.com/sabrusrin/wildberries_st5/pkg/facade"
	"github.com/sabrusrin/wildberries_st5/pkg/mentor"
	"github.com/sabrusrin/wildberries_st5/pkg/models"
	"github.com/sabrusrin/wildberries_st5/pkg/practice"
	"github.com/sabrusrin/wildberries_st5/pkg/theory"
)

var testOkEmpty = `
Read books:
Listen what your mentor tells you
Do the practice tasks given by mentors:
`
var testOkOne = `
Read books:  Go101
Listen what your mentor tells you
Do the practice tasks given by mentors: Implement facade pattern
`

var testOkTwo = `
Read books: Go101, Concurrency in Go
Listen what your mentor tells you
Do the practice tasks given by mentors:  Implement facade pattern, Implement builder pattern
`

func TestOk(t *testing.T)	{
	th := theory.NewBookList(models.TheoryHeader)
	m := mentor.NewMentor(models.MentorHeader)
	p := practice.NewTaskList(models.PracticeHeader)
	wbTrial := facade.NewPlanner(th, m, p)
	// testing for empty string argument, adding a task to the task list and appending task list
	out := wbTrial.Plan("", "")
	if out != testOkEmpty {
		t.Errorf("Test for empty string argument failed!\nExpected:\n%v\nGot %v", testOkEmpty, out)
	}
	out = wbTrial.Plan("Go101", "Implement facade pattern")
	if out != testOkOne	{
		t.Errorf("Test for valid argument failed!\nExpected:\n%v\nGot:\n%v", testOkOne, out)
	}
	out = wbTrial.Plan("Concurrency in Go", "Implement builder pattern")
	if out != testOkTwo	{
		t.Errorf("Test for correct task list appending failed!\nExpected:\n%v\nGot %v", testOkTwo, out)
	}
}
