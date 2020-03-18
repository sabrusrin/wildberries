package facade

type Theory struct	{
}

type Coursera struct {
}

type Practice struct {
}

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

func (th *Theory) Read() string	{
	return "Read books\n"
}

func (c *Coursera) Watch() string	{
	return "Watch <<Разработка веб-сервисов на GO>> course on Coursera\n"
}

func (p *Practice) Solve() string	{
	return "Do the practice tasks given by mentors\n"
}

