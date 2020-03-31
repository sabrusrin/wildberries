package main

import (
	"fmt"

	"github.com/sabrusrin/wildberries_st5/pkg/visitor"
	"github.com/sabrusrin/wildberries_st5/pkg/country"
	"github.com/sabrusrin/wildberries_st5/pkg/city"
)

func main() {
	moscowSights := []string{"Red Square", "Bolshoi Theatre", "Zaryad'ye",}
	kazanSights := []string{"Kremlin", "Kaban lake", "Bauman street"}
	piterSights := []string{"Kazan Cathedral", "Church of the Savior on Spilled Blood", "Winter Palace"}

	tourist := visitor.NewVisitor()

	moscow := city.NewCity("Moscow", moscowSights, tourist)
	kazan := city.NewCity("Kazan", kazanSights, tourist)
	stPete := city.NewCity("St. Petersburg", piterSights, tourist)

	russia := country.NewCountry("Russia", moscow, kazan, stPete)
	sightsToSee := russia.Accept()
	fmt.Printf("%v", sightsToSee)
}
