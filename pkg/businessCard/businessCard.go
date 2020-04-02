package businessCard

import "fmt"

type siteInfo interface {
	Append(string)
	Return() string
}

// Interface to work with site, also a concreteBuilder
type BusinessCardSite interface {
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
	fmt.Print("Enter your Name Surname/Company name: ")
	if _, err = fmt.Scan(&buffer); err == nil {
		s.siteInfo.Append("<body bgcolor=\"F5F5DC\">\n<hr>\n<center>\n<h1>" + buffer + "</h1>\n")
		fmt.Print("Enter your Job title/Business area: ")
		if _, err := fmt.Scan(&buffer); err == nil {
			s.siteInfo.Append("<h2>" + buffer + "</h2>\n")
			fmt.Print("Enter your e-mail address: ")
			if _, err := fmt.Scan(&buffer); err == nil {
				s.siteInfo.Append("E-mail: <a href=\"" + buffer + "\">" + buffer + "</a>\n")
				fmt.Print("Enter your phone number: ")
				if _, err := fmt.Scan(&buffer); err == nil {
					s.siteInfo.Append("<p>Phone number: " + buffer + "\n")
					s.siteInfo.Append("</center>\n<hr>\n</body>\n</html>\n")
				}
			}
		}

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

// NewBusinessCardSite
func NewBusinessCardSite(s siteInfo) BusinessCardSite {
	return &site{
		siteInfo: s,
	}
}
