package theory

// Interface to work with theory type

type Theory struct	{
	book	string
}

func (t *Theory) Read(s string) string	{
	if len(s) != 0	{
		if len(t.book) == 0	{
			t.book = t.book + " " + s
		} else	{
			t.book = t.book + ", " + s
		}
	}
	return "Read books:" + t.book + "\n"
}

// Constructor for theory type
func NewTheory() Theory	{
	return Theory{}
}
