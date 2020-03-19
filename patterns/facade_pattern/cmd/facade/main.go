package main

import (
	"fmt"
	"../../pkg/facade" //don't know how to import my package
)

func main()	{

	in := "Read books\nWatch <<Разработка веб-сервисов на GO>> course on Coursera\n" +
				"Do the practice tasks given by mentors\n"
	wbTrial := InitTrial()

	out := wbTrial.Plan()

	if in != out	{
		fmt.Printf("Expected: %s\nGot: %s\n", in, out)
	}
}