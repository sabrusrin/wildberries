package main

import (
	"fmt"
	"pkg/visitor"
	"pkg/country"
	"pkg/city"
)

func main() {

	moscow := city.NewCity("Moscow", map[string]bool{"Red Square":false,
														"Bolshoi Theatre":false, "Zaryad'ye":false})
	russia := country.NewCountry("Russia")
	russia.Add(moscow)
	tourist := visitor.NewVisitor()
	sightsToSee := russia.Accept(tourist)
}
