package facade

/*
** A structure to represent Facade
 */
type TrialPeriod struct	{
	Theory		*theory
	Coursera	*coursera
	Practice	*practice
}

// Method for a TrialPeriod which returns a plan for trial period
func (t *TrialPeriod) Plan() string	{
	res := t.Theory.read() + t.Coursera.watch() + t.Practice.solve()
	return res
}
