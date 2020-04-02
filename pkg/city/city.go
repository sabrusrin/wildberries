package city

import "fmt"

type visitor interface {
	VisitCity(City) (string, error)
}

// City ...
type City interface {
	Accept() (string, error)
	SightsToSee() (string, error)
}

type city struct {
	cityName   string
	sightsList []string
	visitor
}

func (c *city) Accept() (res string, err error) {
	res, err = c.visitor.VisitCity(c)
	return
}

// SightsToSee returns a list of sights to see in the city
func (c *city) SightsToSee() (res string, err error) {
	var toSee string
	err = nil
	if len(c.sightsList) != 0 {
		toSee = "In " + c.cityName + " you should check following sights: "
		for i, sight := range c.sightsList {
			if i+1 < len(c.sightsList) {
				toSee += sight + ", "
			} else {
				toSee += sight + "\n"
			}
		}
	} else {
		err = fmt.Errorf("error. Sight list for city %s is Empty", c.cityName)
	}
	res = toSee
	return
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
