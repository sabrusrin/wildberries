package businessCard

import "fmt"

// Interface to work with site, also a concreteBuilder
type BusinessCardSite interface {
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
	s.siteInfo.Append("<head>\n<title>" + buffer + "</title>\n</head>\n")
}

// BuildSiteBody will ask for certain parameters and append them to the site complex object
func (s *site) BuildSiteBody()	{
	var buffer string
	fmt.Print("Enter your Name Surname/Company name: ")
	fmt.Scan(&buffer)
	s.siteInfo.Append("<body bgcolor=\"F5F5DC\">\n<hr>\n<center>\n<h1>" + buffer + "</h1>\n")
	fmt.Print("Enter your Job title/Business area: ")
	fmt.Scan(&buffer)
	s.siteInfo.Append("<h2>" + buffer + "</h2>\n")
	fmt.Print("Enter your e-mail address: ")
	fmt.Scan(&buffer)
	s.siteInfo.Append("E-mail: <a href=\"" + buffer + "\">" + buffer + "</a>\n")
	fmt.Print("Enter your phone number: ")
	fmt.Scan(&buffer)
	s.siteInfo.Append("<p>Phone number: " + buffer + "\n")
	s.siteInfo.Append("</center>\n<hr>\n</body>\n</html>")
}

// ReturnSite returns a site product
func (s *site) ReturnSite() string {
	return s.siteInfo.Return()
}

// NewBusinessCardSite
func NewBusinessCardSite(s siteInfo) BusinessCardSite	{
	return &site{
		siteInfo: s,
	}
}
