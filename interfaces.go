package staticIntf

import (
	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/htmlDoc"
)

type PageDto interface {
	Id() int
	Title() string
	TitlePlain() string
	ThumbUrl() string
	ImageUrl() string
	Description() string
	DisqusId() string
	CreateDate() string
	Content() string
	Category() string
	Url() string
	Domain() string
	PathFromDocRoot() string
	FsPath() string
	HtmlFilename() string
	ThumbBase64() string
}

type ConfigContainer interface {
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
	HomeText() string
	HomeHeadline() string
}

type LocationContainer interface {
	AddMain(Location)
	Main() []Location
	AddMarginal(Location)
	Marginal() []Location
}

type PagesContainer interface {
	Pages() []Page
	NaviPages() []Page
	Representationals() []Page

	AddPage(Page)
	AddRepresentational(Page)
	AddNaviPage(Page)

	Variant() string
}

type PagesContainerCollection interface {
	AddContainer(PagesContainer)
	Containers() []PagesContainer
	ContainersOrderedByVariants(...string) []PagesContainer
	Posts() []Page
	Portfolio() []Page
	Home() []Page
	PostNaviPages() []Page
	Marginals() []Page
	Narratives() []Page
	NarrativeMarginals() []Page
	Pages() []Page
}

type Site interface {
	ConfigContainer
	LocationContainer
	PagesContainerCollection
}

type Context interface {
	RenderPages() []fs.FileContainer
	GetComponents() []Component
}

type Renderer interface {
	AddComponents(...Component)
	Components() []Component
	Site() Site

	Pages(...Page) []Page
	AddPage(Page)

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
	FooterNavigationLocations() []Location
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
	Id() int
	PublishedTime(...string) string
	Description() string
	Content() string
	ImageUrl() string
	DisqusId() string
	ThumbBase64() string
	Category() string
	GetDoc() *htmlDoc.HtmlDoc
}

type Page interface {
	Location
	PageContent
	AcceptVisitor(v Component)
	AddBodyNodes([]*htmlDoc.Node)
	AddHeaderNodes([]*htmlDoc.Node)
	NavigatedPages(...Page) []Page
}

type Component interface {
	VisitPage(Page)
	GetCss() string
	GetJs() string
	Renderer(...Renderer) Renderer
}
