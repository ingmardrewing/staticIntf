package staticIntf

import (
	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/htmlDoc"
)

type ContextGroup interface {
	RenderPages(string) []fs.FileContainer
	GetComponents() []Component
}

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
	GetDisqusShortname() string
	GetMainNavigationLocations() []Location
	GetReadNavigationLocations() []Location
	GetFooterNavigationLocations() []Location
	GetElements() []Page
	SetElements([]Page)
	SetContextDto(ContextDto)
	FsSetOff(...string) string
	AddRss()
	AddComponent(c Component)
	GetComponents() []Component
	RenderPages(targetDir string) []fs.FileContainer
	AddPage(p Page)
}

type Location interface {
	Domain(...string) string
	Title(...string) string
	ThumbnailUrl(...string) string
	Url(...string) string
	PathFromDocRoot(...string) string
	HtmlFilename(...string) string
}

type PageContent interface {
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

type ContextDto interface {
	TwitterHandle(...string) string
	Topic(...string) string
	Tags(...string) string
	Site(...string) string
	CardType(...string) string
	Section(...string) string
	FBPage(...string) string
	TwitterPage(...string) string
	Rss(...string) string
	Css(...string) string
	Domain(...string) string
	DisqusId(...string) string
	TargetDir(...string) string
}

type Page interface {
	Location
	PageContent
}

type NaviPage interface {
	Location
	PageContent
	NavigatedPages(...Page) []Page
}

type Component interface {
	VisitPage(p Page)
	GetCss() string
	GetJs() string
	SetContext(context Context)
}
