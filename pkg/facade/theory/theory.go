package theory

// Interface to work with theory type
type Reader interface {
	Read() string
}
type theory struct	{
}

func (t *theory) Read() string	{
	return "Read books\n"
}

// Constructor for theory type
func NewTheory() Reader	{
	return &theory{}
}
