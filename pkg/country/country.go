package country

type visitor interface {
}

// Place provides an interface for place that the visitor should visit.
type place interface {
	Accept(v visitor) []string
}

type Country interface {
	Add(p place)
	Accept(v visitor) [][]string
}

// City implements a collection of places to visit.
type country struct {
	countryName string
	places      []place
}

// Add appends Place to the collection.
func (c *country) Add(p place) {
	c.places = append(c.places, p)
}

// Accept implements a visit to all places in the city.
func (c *country) Accept(v visitor) [][]string {
	var result [][]string
	for _, p := range c.places {
		result = append(result, p.Accept(v))
	}
	return result
}

func NewCountry(name string) Country {
	return &country{
		countryName: name,
	}
}
