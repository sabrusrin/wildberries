package product

// Interface to work with complex object
type Site interface {
	Append(string)
	Return() string
}

type site struct {
	site     string
	siteType string
}

// Append appends to a site
func (s *site) Append(str string) {
	s.site += str
}

// Return returns a site
func (s *site) Return() string {
	return s.site
}

// NewSite ...
func NewSite(s string) Site {
	return &site{
		site:     "",
		siteType: s,
	}
}
