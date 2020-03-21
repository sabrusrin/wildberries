package theory

// Interface to work with theory type

type Theory struct	{
	book	string
}

func (t *Theory) Read(s string) string	{
	t.book = t.book + " " + s + ","
	return "Read books:" + t.book + "\n"
}

// Constructor for theory type
func NewTheory() Theory	{
	return Theory{}
}
