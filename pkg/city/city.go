package city

import "fmt"

type visitor interface {
	VisitCity(*city) []string
}

type City interface {
	Accept(v visitor) []string
//	Return() (map[string]bool, string)
}

type city struct {
	cityName   string
	sightsList map[string]bool
}

func (c *city) Accept(v visitor) []string {
	return v.VisitCity(c)
}

func (c *city) Return() (map[string]bool, string) {
	return c.sightsList, c.cityName
}

func (c *city) SightsToSee() []string {
	var sight string
	var toSee []string
	var err error
	fmt.Printf("What sights you've already seen in %s?(press 0 when you'll recap)\n", c.cityName)
	for _, err = fmt.Scan(&sight); sight != "0" && err == nil; {
		if _, ok := c.sightsList[sight]; ok {
			c.sightsList[sight] = true
		}
	}
	if err != nil {
		panic(err)
	}
	for sight, status := range c.sightsList {
		if status == false {
			toSee = append(toSee, sight)
		}
	}
	return toSee
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
