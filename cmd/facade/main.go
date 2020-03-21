package main

import (
	"WB/pkg/facade"
	"fmt"
)

func main()	{
	in := "Read books: Go101,\n" + "Listen what your mentor tells you\n" +
			"Do the practice tasks given by mentors: Implement facade pattern,\n"

	wbTrial := facade.NewTrial()

	// Adding a new book and task to the trialPeriodPlan and comparing with expected Plan
	out := wbTrial.Plan("Go101", "Implement facade pattern")
	if in != out	{
		fmt.Printf("Expected: %s\nGot: %s\n", in, out)
	} else	{
		fmt.Printf("%sYou are doing fine!\n", out)
	}

	// Adding new books and tasks to the trialPeriodPlan
	out = wbTrial.Plan("Concurrency in go", "Implement builder pattern")
	fmt.Printf("\n%s", out)
}