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
	var err error
	fmt.Print("Enter name for new file: ")
	if _, err = fmt.Scan(&buffer); err == nil {
		fileName := buffer + ".html"
		f, err := os.Create(fileName)
		if err == nil {
			defer f.Close()
			fmt.Print("What type of site you want to build(BusinessCard or Photogallery): ")
			if _, err = fmt.Scan(&buffer); err == nil {
				site := product.NewSite(buffer)
				var siteBuilder businessCard.BusinessCardSite
				if buffer == "BusinessCard" {
					siteBuilder = businessCard.NewBusinessCardSite(site)
				} else if buffer == "Photogallery" {
					siteBuilder = photogallery.NewPhotogallerySite(site)
				} else {
					err = fmt.Errorf("unknown type of site: %s", buffer)
					panic(err)
				}
				director := builder.NewDirector(siteBuilder)
				if out, err := director.BuildSite(); err == nil {
					if _, err = f.WriteString(out); err != nil {
						fmt.Printf("Your site written to: %s", fileName)
					}
				}
			}
		}
	}
	if err != nil {
		panic(err)
	}
}
