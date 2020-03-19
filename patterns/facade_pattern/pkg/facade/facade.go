package facade

/*
** A structure to represent Facade
 */
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

// Method for a TrialPeriod which returns a plan for trial period
func (t *TrialPeriod) Plan() string	{
	res := t.theory.Read() + t.coursera.Watch() + t.practice.Solve()
	return res
}

type Theory struct	{
}

// Method to work with Theory struct
func (th *Theory) Read() string	{
	return "Read books\n"
}

type Coursera struct {
}

// Method to work with Coursera struct
func (c *Coursera) Watch() string	{
	return "Watch <<Разработка веб-сервисов на GO>> course on Coursera\n"
}

type Practice struct {
}

// Method to work with Practice struct
func (p *Practice) Solve() string	{
	return "Do the practice tasks given by mentors\n"
}

