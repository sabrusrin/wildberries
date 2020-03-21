package mentor

// Interface to work with mentor type
type Mentor struct	{
}

func (c *Mentor) Listen() string	{
	return "Listen what your mentor tells you\n"
}

// Constructor for mentor type
func NewMentor() Mentor	{
	return Mentor{}
}
