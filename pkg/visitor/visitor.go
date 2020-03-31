package visitor

import (
	"github.com/sabrusrin/wildberries_st5/pkg/city"
)

// Visitor provides a visitor interface.
type Visitor interface {
	VisitCity(city.City) string
}

type people struct {
}

// VisitCity ...
func (v *people) VisitCity(city city.City) string {
	return city.SightsToSee()
}

// NewVisitor ...
func NewVisitor() Visitor {
	return &people{}
}
