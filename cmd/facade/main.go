package main

import (
	"pkg/facade"
	"pkg/mentor"
	"pkg/models"
	"pkg/practice"
	"pkg/theory"
	"fmt"
)

func main()	{
	in := "Read books: Go101\n" + "Listen what your mentor tells you\n" +
			"Do the practice tasks given by mentors: Implement facade pattern\n"

	t := theory.NewTheory(models.TheoryHeader)
	m := mentor.NewMentor(models.MentorHeader)
	p := practice.NewPractice(models.PracticeHeader)
	wbTrial := facade.NewPlanner(&t, &m, &p)
	// Adding a new book and task to the trialPeriodPlan and comparing with expected Plan
	out := wbTrial.Plan("Go101", "Implement facade pattern")
	if in != out	{
		fmt.Printf("Expected: %s\nGot: %s\n", in, out)
	} else	{
		fmt.Printf("%sYou are doing fine!\n", out)
	}
	// Appending the task list with new books and tasks
	out = wbTrial.Plan("Concurrency in go", "Implement builder pattern")
	fmt.Printf("\n%s", out)
}
