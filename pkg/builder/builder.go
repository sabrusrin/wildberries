package builder

type siteBuilder interface {
	StartNewSite() error
	BuildSiteBody() error
	ReturnSite() (string, error)
}

// Director interface to work with builder
type Director interface {
	BuildSite() (string, error)
}

type director struct {
	builder siteBuilder
}

// BuildSite calls necessary commands to build a site
func (d *director) BuildSite() (res string, err error) {
	if err = d.builder.StartNewSite(); err == nil {
		if err = d.builder.BuildSiteBody(); err == nil {
			res, err = d.builder.ReturnSite()
		}
	}
	return
}

// NewDirector ...
func NewDirector(s siteBuilder) Director {
	return &director{
		builder: s,
	}
}
