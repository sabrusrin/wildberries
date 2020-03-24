package photogallery

import "fmt"

// Interface to work with site, also a concreteBuilder
type PhotogallerySite interface {
	StartNewSite()
	BuildSiteBody()
	ReturnSite() string
}

type siteInfo interface {
	Append(string)
	Return() string
}

type site struct {
	siteInfo siteInfo
}

// StartNewSite asks for a site title and appends it to the site complex object
func (s *site) StartNewSite()	{
	var buffer string
	fmt.Print("Enter your site title: ")
	fmt.Scan(&buffer)
	s.siteInfo.Append("<!DOCTYPE html>\n<html>\n<head>\n<title>" + buffer + "</title>\n</head>\n")
}

// BuildSiteBody will ask for certain parameters and append them to the site complex object
func (s *site) BuildSiteBody()	{
	var buffer string
	var num int
	fmt.Print("Enter photogallery name: ")
	fmt.Scan(&buffer)
	s.siteInfo.Append("<body bgcolor=\"87cefa\">\n<hr>\n<center>\n<h1>" + buffer + "</h1>\n<p>")
	fmt.Print("How many photos you want to show: ")
	fmt.Scan(&num)
	for num != 0	{
		fmt.Print("Enter path to your image: ")
		fmt.Scan(&buffer)
		s.siteInfo.Append("<img src=\"" + buffer + "\"align=\"bottom\">\n")
		num--
	}
	s.siteInfo.Append("</center>\n<hr>\n</body>\n</html>")
}

// ReturnSite returns a site product
func (s *site) ReturnSite() string {
	return s.siteInfo.Return()
}

// NewPhotogallerySite ...
func NewPhotogallerySite(s siteInfo) PhotogallerySite	{
	return &site{
		siteInfo: s,
	}
}
