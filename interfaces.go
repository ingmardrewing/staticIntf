package staticIntf

import (
	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/htmlDoc"
)

type Context interface {
	GetTwitterHandle() string
	GetContentSection() string
	GetContentTags() string
	GetSiteName() string
	GetTwitterCardType() string
	GetOGType() string
	GetFBPageUrl() string
	GetTwitterPage() string
	GetCssUrl() string
	GetCss() string
	GetRssUrl() string
	GetHomeUrl() string
	GetDisqusShortname() string
	GetMainNavigationLocations() []Location
	GetReadNavigationLocations() []Location
	GetFooterNavigationLocations() []Location
	GetElements() []Page
	SetElements([]Page)
	AddComponent(c Component)
	GetComponents() []Component
	RenderPages(targetDir string) []fs.FileContainer
	AddPage(p Page)
	SetGlobalFields(twitterHandle, topic, tags, site, cardType, section, fbPage, twitterPage, cssUrl, rssUrl, home, disqusShortname string)
}

type Location interface {
	Domain(...string) string
	Title(...string) string
	ThumbnailUrl(...string) string
	Url(...string) string
	PathFromDocRoot(...string) string
	HtmlFilename(...string) string
}

type Page interface {
	Location
	Id(...int) int
	AcceptVisitor(v Component)
	AddBodyNodes([]*htmlDoc.Node)
	AddHeaderNodes([]*htmlDoc.Node)

	PublishedTime(...string) string
	Description(...string) string
	Content(...string) string
	ImageUrl(...string) string
	DisqusId(...string) string
	GetDoc() *htmlDoc.HtmlDoc
}

type Component interface {
	VisitPage(p Page)
	GetCss() string
	GetJs() string
	SetContext(context Context)
}
