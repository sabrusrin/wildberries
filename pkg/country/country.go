package country

import (
	"github.com/sabrusrin/wildberries_st5/pkg/city"
)

type (
	cities = []city.City
)

// Country ...
type Country interface {
	Accept() ([]string, error)
}

type country struct {
	countryName string
	cities
}

// Accept implements a visit to all cities in the .
func (c *country) Accept() (res []string, err error) {
	var tmp string
	err = nil
	for _, p := range c.cities {
		if tmp, err = p.Accept(); err != nil {
			break
		}
		res = append(res, tmp)
	}
	return
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
