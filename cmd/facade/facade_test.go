package main

import (
	"WB/pkg/facade"
	"WB/pkg/mentor"
	"WB/pkg/models"
	"WB/pkg/practice"
	"WB/pkg/theory"
	"testing"
)

func TestOk(t *testing.T)	{
	th := theory.NewTheory(models.TheoryHeader)
	m := mentor.NewMentor(models.MentorHeader)
	p := practice.NewPractice(models.PracticeHeader)
	wbTrial := facade.NewPlanner(&th, &m, &p)

	//testing for empty string argument, adding a task to the task list and appending task list
	okResult := "Read books:\n" + "Listen what your mentor tells you\n" +
		"Do the practice tasks given by mentors:\n"
	out := wbTrial.Plan("", "")
	if out != okResult	{
		t.Errorf("Test for empty string argument failed!\nExpected:\n%v\nGot %v", okResult, out)
	}
	okResult = "Read books: Go101\n" + "Listen what your mentor tells you\n" +
		"Do the practice tasks given by mentors: Implement facade pattern\n"
	out = wbTrial.Plan("Go101", "Implement facade pattern")
	if out != okResult	{
		t.Errorf("Test for valid argument failed!\nExpected:\n%v\nGot:\n%v", okResult, out)
	}
	okResult = "Read books: Go101, Concurrency in Go\n" + "Listen what your mentor tells you\n" +
		"Do the practice tasks given by mentors: Implement facade pattern, Implement builder pattern\n"
	out = wbTrial.Plan("Concurrency in Go", "Implement builder pattern")
	if out != okResult	{
		t.Errorf("Test for correct task list appending failed!\nExpected:\n%v\nGot %v", okResult, out)
	}
}
