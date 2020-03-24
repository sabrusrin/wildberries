package builder

type siteBuilder interface {
	StartNewSite()
	BuildSiteBody()
	ReturnSite() string
}

// Director interface to work with builder
type Director interface {
	BuildSite() string
}

type director struct {
	builder siteBuilder
}

// BuildSite calls necessary commands to build a site
func (d *director) BuildSite() string {
	d.builder.StartNewSite()
	d.builder.BuildSiteBody()
	return d.builder.ReturnSite()
}

// NewDirector ...
func NewDirector(s siteBuilder) Director {
	return &director{
		builder: s,
	}
}
