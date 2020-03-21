package main

import (
	"WB/pkg/facade"
	"fmt"
)

func main()	{
	in := "Read books\nListen what your mentor tells you\n" +
				"Do the practice tasks given by mentors\n"

	wbTrial := facade.NewTrial()
	out := wbTrial.Plan()
	if in != out	{
		fmt.Printf("Expected: %s\nGot: %s\n", in, out)
	}
}