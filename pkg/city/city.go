package city

type visitor interface {
	VisitCity(*city) []string
}

type City interface {
	Accept(v visitor) []string
	Return() (map[string]bool, string)
}

type city struct {
	cityName string
	sightsList   map[string]bool
}

func (c *city) Accept(v visitor) []string {
	return v.VisitCity(c)
}

func (c *city) Return() (map[string]bool, string) {
	return c.sightsList, c.cityName
}


func NewCity(
	name string,
	sights map[string]bool,
	) City {
	return &city{
		cityName:   name,
		sightsList: sights,
	}
}
