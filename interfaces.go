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
	Pages() []Page
	Marginals() []Page
	Narratives() []Page

	TwitterHandle() string
	Topic() string
	Tags() string
	Site() string
	CardType() string
	Section() string
	FBPage() string
	TwitterPage() string
	Rss() string
	Css() string
	Domain() string
	DisqusId() string
	TargetDir() string
}

type ContextGroup interface {
	RenderPages() []fs.FileContainer
	GetComponents() []Component
}

type Context interface {
	AddComponents(...Component)
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
	GetDisqusShortname() string
	GetMainNavigationLocations() []Location
	GetReadNavigationLocations() []Location
	GetFooterNavigationLocations() []Location
	GetElements() []Page
	SetElements([]Page)
	FsSetOff(...string) string
	AddComponent(c Component)
	GetComponents() []Component
	RenderPages() []fs.FileContainer
	GetPages() []Page
	AddPage(p Page)
	SiteDto() Site
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
	VisitPage(p Page)
	GetCss() string
	GetJs() string
	SetContext(context Context)
}
