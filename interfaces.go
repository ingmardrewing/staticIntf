package staticIntf

import (
	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/htmlDoc"
)

type Site interface {
	AddMain(Location)
	Main() []Location
	AddMarginal(Location)
	Marginal() []Location

	Posts() []Page
	PostNaviPages() []Page
	Pages() []Page
	Marginals() []Page
	Narratives() []Page
	NarrativeMarginals() []Page

	TwitterHandle() string
	Topic() string
	Tags() string
	Site() string
	CardType() string
	Section() string
	FBPage() string
	TwitterPage() string
	RssPath() string
	RssFilename() string
	Css() string
	Domain() string
	DisqusId() string
	TargetDir() string
}

type Context interface {
	RenderPages() []fs.FileContainer
	GetComponents() []Component
}

type Renderer interface {
	AddComponents(...Component)
	Components() []Component

	Site() Site
	TwitterHandle() string
	ContentSection() string
	ContentTags() string
	SiteName() string
	TwitterCardType() string
	OGType() string
	FBPageUrl() string
	TwitterPage() string
	CssUrl() string
	Css() string
	DisqusShortname() string
	MainNavigationLocations() []Location
	ReadNavigationLocations() []Location
	FooterNavigationLocations() []Location
	FsSetOff(...string) string
	Pages(...Page) []Page
	AddPage(Page)

	Render() []fs.FileContainer
}

type Location interface {
	Domain(...string) string
	PathFromDocRoot(...string) string
	HtmlFilename(...string) string

	Title(...string) string
	ThumbnailUrl(...string) string
	ExternalLink(...string) string

	// Complete Url including protocol, domain, port (if any),
	// path from docroot and filename.
	// Derived from other fields of Location.
	Url() string

	// Compose path from doc root including
	// the filename.
	PathFromDocRootWithName() string
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
	ThumbBase64(...string) string
	Category(...string) string
	GetDoc() *htmlDoc.HtmlDoc
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
	VisitPage(Page)
	GetCss() string
	GetJs() string
	Renderer(...Renderer) Renderer
}
