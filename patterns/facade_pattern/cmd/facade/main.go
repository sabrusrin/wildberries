package main

import (
	"fmt"
	"strings"
)

func main()	{

	in := "Read books\nWatch <<Разработка веб-сервисов на GO>> course on Coursera\n" +
				"Do the practice tasks given by mentors\n"
	wbTrial := initTrial()

	out := wbTrial.Plan()

	if (in != out)	{
		fmt.Errorf("Expected: %s\nGot: %s\n", in, out)
	}
}