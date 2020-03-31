package country

import (
	"pkg/city"
)

type (
	cities = []city.City
)

// Country ...
type Country interface {
	Accept() []string
}

type country struct {
	countryName string
	cities
}

// Accept implements a visit to all cities in the .
func (c *country) Accept() []string {
	var result []string
	for _, p := range c.cities {
		result = append(result, p.Accept())
	}
	return result
}

// NewCountry ...
func NewCountry(
	name string,
	c ...city.City,
) Country {
	var allCities cities
	for _, c := range c {
		allCities = append(allCities, c)
	}
	return &country{
		countryName: name,
		cities:      allCities,
	}
}
