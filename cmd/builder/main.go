package main

import (
	"fmt"
	"os"

	"github.com/sabrusrin/wildberries_st5/pkg/builder"
	"github.com/sabrusrin/wildberries_st5/pkg/businessCard"
	"github.com/sabrusrin/wildberries_st5/pkg/photogallery"
	"github.com/sabrusrin/wildberries_st5/pkg/product"
)

func main() {
	var buffer string
	fmt.Print("Enter name for new file: ")
	fmt.Scan(&buffer)
	fileName := buffer + ".html"
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Print("What type of site you want to build(BusinessCard or Photogallery): ")
	fmt.Scan(&buffer)
	site := product.NewSite(buffer)
	var siteBuilder businessCard.BusinessCardSite
	if buffer == "BusinessCard" {
		siteBuilder = businessCard.NewBusinessCardSite(site)
	} else if buffer == "Photogallery" {
		siteBuilder = photogallery.NewPhotogallerySite(site)
	} else {
		fmt.Errorf("Unknown type of site: %s", buffer)
	}
	director := builder.NewDirector(siteBuilder)
	out := director.BuildSite()

	f.WriteString(out)
	fmt.Printf("Your site written to: %s", fileName)
}
