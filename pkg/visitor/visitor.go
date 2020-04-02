package visitor

import (
	"github.com/sabrusrin/wildberries_st5/pkg/city"
)

// Visitor provides a visitor interface.
type Visitor interface {
	VisitCity(city.City) (string, error)
}

type people struct {
}

// VisitCity ...
func (v *people) VisitCity(city city.City) (res string, err error) {
	res, err = city.SightsToSee()
	return
}

// NewVisitor ...
func NewVisitor() Visitor {
	return &people{}
}
