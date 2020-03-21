package theory

// Interface to work with theory type
type Reader interface {
	Read(string) string
}
type theory struct	{
	book	string
}

func (t *theory) Read(s string) string	{
	t.book = t.book + " " + s + ","
	return "Read books:" + t.book + "\n"
}

// Constructor for theory type
func NewTheory() Reader	{
	return &theory{}
}
