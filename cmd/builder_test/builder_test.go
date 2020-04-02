package builder_test

import (
	"testing"

	"github.com/sabrusrin/wildberries_st5/pkg/builder"
	"github.com/sabrusrin/wildberries_st5/pkg/businessCard"
	"github.com/sabrusrin/wildberries_st5/pkg/photogallery"
	"github.com/sabrusrin/wildberries_st5/pkg/product"
)

var TestOkBusinessCard = `<!DOCTYPE html>
<html>
<head>
<title></title>
</head>
<body bgcolor="F5F5DC">
<hr>
<center>
<h1></h1>
<h2></h2>
E-mail: <a href=""></a>
<p>Phone number: 
</center>
<hr>
</body>
</html>
`

var TestOkPhotogallery = `<!DOCTYPE html>
<html>
<head>
<title></title>
</head>
<body bgcolor="87cefa">
<hr>
<center>
<h1></h1>
<p></center>
<hr>
</body>
</html>
`

var testFail = `no site was created! check function call order`

func TestOk(t *testing.T) {
	site := product.NewSite("BusinessCard")
	siteBuilder := businessCard.NewBusinessCardSite(site)
	director := builder.NewDirector(siteBuilder)
	out, err := director.BuildSite()
	if out != TestOkBusinessCard && err == nil {
		t.Errorf("Test for BusinessCard site failed!\nExpected:\n%v", TestOkBusinessCard)
	}

	site = product.NewSite("Photogallery")
	siteBuilder = photogallery.NewPhotogallerySite(site)
	director = builder.NewDirector(siteBuilder)
	out, err = director.BuildSite()
	if out != TestOkPhotogallery && err == nil {
		t.Errorf("Test for Photogallery site failed!\nExpected:\n%v", TestOkPhotogallery)
	}
}

func TestForError(t *testing.T) {
	site := product.NewSite("BusinessCard")
	siteBuilder := businessCard.NewBusinessCardSite(site)
	_, err := siteBuilder.ReturnSite()
	if err == nil {
		t.Errorf("Incorrect error handling!\nExpected:%s", testFail)
	}
	site = product.NewSite("Photogallery")
	siteBuilder = photogallery.NewPhotogallerySite(site)
	_, err = siteBuilder.ReturnSite()
	if err == nil {
		t.Errorf("Incorrect error handling!\nExpected:%s", testFail)
	}
}
