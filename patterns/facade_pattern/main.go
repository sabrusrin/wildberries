package main

import (
	"fmt"
)

type TrialPeriod struct	{
	theory		*Theory
	coursera	*Coursera
	practice	*Practice
}

func InitTrial() *TrialPeriod	{
	return &TrialPeriod{
		theory: &Theory{},
		coursera: &Coursera{},
		practice: &Practice{},
	}
}

func (t *TrialPeriod) Plan() string	{
	res := t.theory.Read() + t.coursera.Watch() + t.practice.Solve()

	return res
}

type Theory struct	{
}

func (th *Theory) Read() string	{
	return "Read books\n"
}

type Coursera struct {
}

func (c *Coursera) Watch() string	{
	return "Watch <<Разработка веб-сервисов на GO>> course on Coursera\n"
}

type Practice struct {
}

func (p *Practice) Solve() string	{
	return "Do the practice tasks given by mentors\n"
}

func main()	{

	in := "Read books\nWatch <<Разработка веб-сервисов на GO>> course on Coursera\n" +
				"Do the practice tasks given by mentors\n"
	wbTrial := InitTrial()

	out := wbTrial.Plan()

	if in != out	{
		fmt.Printf("Expected: %s\nGot: %s\n", in, out)
	}
}