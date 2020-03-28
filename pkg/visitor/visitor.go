package visitor

import (
	"fmt"
)

type place interface {
	Return() (map[string]bool, string)
}

// Visitor provides a visitor interface.
type Visitor interface {
	VisitCity(p place) []string
}

// People implements the Visitor interface.
type people struct {
}

func (v *people) VisitCity(p place) []string {
	sightList, cityName := p.Return()
	return SightsToSee(sightList, cityName)
}

func SightsToSee(sightList map[string]bool, cityName string) []string {
	var sight string
	var toSee []string
	var err   error
	fmt.Printf("What sights you've already seen in %s?(press 0 when you'll recap)\n", cityName)
	for _, err = fmt.Scan(&sight); sight != "0" && err == nil {
		if _, ok := sightList[sight]; ok {
			sightList[sight] = true
		}
	}
	if err != nil {
		panic(err)
	}
	for sight, status := range sightList {
		if status == false {
			toSee = append(toSee, sight)
		}
	}
	return toSee
}

func NewVisitor() Visitor {
	return &people{}
}
