package mentor

// Interface to work with mentor type
type Listener interface {
	Listen() string
}

type mentor struct	{
}

func (c *mentor) Listen() string	{
	return "Listen what your mentor tells you\n"
}

// Constructor for mentor type
func NewMentor() Listener	{
	return &mentor{}
}
