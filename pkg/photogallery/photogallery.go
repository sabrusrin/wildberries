package photogallery

import "fmt"

type siteInfo interface {
	Append(string)
	Return() string
}

// Interface to work with site, also a concreteBuilder
type PhotogallerySite interface {
	StartNewSite() error
	BuildSiteBody() error
	ReturnSite() (string, error)
}

type site struct {
	siteInfo siteInfo
}

// StartNewSite asks for a site title and appends it to the site complex object
func (s *site) StartNewSite() (err error) {
	var buffer string
	fmt.Print("Enter your site title: ")
	if _, err = fmt.Scan(&buffer); err == nil {
		s.siteInfo.Append("<!DOCTYPE html>\n<html>\n<head>\n<title>" + buffer + "</title>\n</head>\n")
	}
	return
}

// BuildSiteBody will ask for certain parameters and append them to the site complex object
func (s *site) BuildSiteBody() (err error) {
	var buffer string
	var num int
	fmt.Print("Enter photogallery name: ")
	if _, err = fmt.Scan(&buffer); err == nil {
		s.siteInfo.Append("<body bgcolor=\"87cefa\">\n<hr>\n<center>\n<h1>" + buffer + "</h1>\n<p>")
		fmt.Print("How many photos you want to show: ")
		if _, err = fmt.Scan(&buffer); err == nil {
			for num != 0 {
				fmt.Print("Enter path to your image: ")
				if _, err = fmt.Scan(&buffer); err == nil {
					s.siteInfo.Append("<img src=\"" + buffer + "\"align=\"bottom\">\n")
					num--
				}
			}
		}
		s.siteInfo.Append("</center>\n<hr>\n</body>\n</html>\n")
	}
	return
}

// ReturnSite returns a site product
func (s *site) ReturnSite() (res string, err error) {
	if res = s.siteInfo.Return(); len(res) == 0 {
		err = fmt.Errorf("no site was created! check function call order")
		return
	}
	err = nil
	return
}

// NewPhotogallerySite ...
func NewPhotogallerySite(s siteInfo) PhotogallerySite {
	return &site{
		siteInfo: s,
	}
}
