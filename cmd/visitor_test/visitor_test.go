package visitor_test

import (
	"testing"

	"github.com/sabrusrin/wildberries_st5/pkg/city"
	"github.com/sabrusrin/wildberries_st5/pkg/country"
	"github.com/sabrusrin/wildberries_st5/pkg/visitor"
)

var testOk = `In Moscow you should check following sights: Red Square, Bolshoi Theatre, Zaryad'ye
In Kazan you should check following sights: Kremlin, Kaban lake, Bauman street
In St. Petersburg you should check following sights: Kazan Cathedral, Church of the Savior on Spilled Blood, Winter Palace
`

func TestOk(t *testing.T) {
	var result string

	moscowSights := []string{"Red Square", "Bolshoi Theatre", "Zaryad'ye",}
	kazanSights := []string{"Kremlin", "Kaban lake", "Bauman street"}
	piterSights := []string{"Kazan Cathedral", "Church of the Savior on Spilled Blood", "Winter Palace"}

	tourist := visitor.NewVisitor()

	moscow := city.NewCity("Moscow", moscowSights, tourist)
	kazan := city.NewCity("Kazan", kazanSights, tourist)
	stPete := city.NewCity("St. Petersburg", piterSights, tourist)

	russia := country.NewCountry("Russia", moscow, kazan, stPete)
	sightsToSee := russia.Accept()

	for _, list := range sightsToSee {
		result += list
	}
	if result != testOk {
		t.Errorf("TestOk failed!\nexpected:\n%s\ngot:\n%s", testOk, result)
	}
}
