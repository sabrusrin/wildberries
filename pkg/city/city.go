package city

type visitor interface {
	VisitCity(City) string
}

// City ...
type City interface {
	Accept() string
	SightsToSee() string
}

type city struct {
	cityName   string
	sightsList []string
	visitor
}

func (c *city) Accept() string {
	return c.visitor.VisitCity(c)
}

// SightsToSee returns a list of sights to see in the city
func (c *city) SightsToSee() string {
	var toSee string
	toSee = "In " + c.cityName + " you should check following sights: "
	for i, sight := range c.sightsList {
		if i+1 < len(c.sightsList) {
			toSee += sight + ", "
		} else {
			toSee += sight + "\n"
		}
	}
	return toSee
}

// NewCity ...
func NewCity(
	name string,
	sights []string,
	visitor visitor,
) City {
	return &city{
		cityName:   name,
		sightsList: sights,
		visitor:    visitor,
	}
}
