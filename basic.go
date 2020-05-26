package gouielement

import (
	"path"
	"strconv"

	"github.com/goog-lukemc/gouidom"
)

type ElementLib struct {
	v *gouidom.VDOM
}

type ElementCFG struct {
	ID                 string
	Parent             string
	Typ                string
	InitializationText string
	Class              []string
	CustomAttributes   map[string]string
}

// NewElementLib Create a new lib for element use
func NewElementLib(v *gouidom.VDOM) *ElementLib {
	return &ElementLib{v: v}
}

// Span create a new span element in the dom with the provided text.
func (c *ElementLib) Span(parent string, text string, classes ...string) *gouidom.Element {
	return c.newElement(&ElementCFG{
		Parent:             parent,
		Typ:                gouidom.HTMLTag.Span,
		InitializationText: text,
		Class:              classes,
	})
}

// WrapperDiv create a new containing div for the elements
func (c *ElementLib) WrapperDiv(parent string) *gouidom.Element {
	return c.newElement(&ElementCFG{
		Parent: parent,
		Typ:    gouidom.HTMLTag.Div,
		Class:  []string{"grid", "container"},
	})
}

func (c *ElementLib) newElement(cfg *ElementCFG) *gouidom.Element {
	ele, err := gouidom.NewElement(cfg.ID, cfg.Parent, cfg.Typ, cfg.InitializationText, cfg.Class...)
	if err != nil {
		gouidom.CLog("%s", err.Error())
	}
	for k, v := range cfg.CustomAttributes {
		ele.SetAttribute(k, v)
	}
	c.v.AddElement(ele)
	return ele
}

// PathOf returns the path of the element for appending
func PathOf(ele *gouidom.Element) string {
	return path.Join(ele.Parent, ele.ID)
}

// InjectStyle injects a new style tag into the body of the document
func (c *ElementLib) InjectStyle(text string) *gouidom.Element {
	return c.newElement(&ElementCFG{
		Parent:             "html/body",
		Typ:                gouidom.HTMLTag.Style,
		InitializationText: text,
	})
}

// Article Adds and article tag usually used to display documentation.
func (c *ElementLib) Article(parent string, classes ...string) *gouidom.Element {
	return c.newElement(&ElementCFG{
		Parent: parent,
		Typ:    gouidom.HTMLTag.Article,
		Class:  classes,
	})
}

// Section Adds and article tag usually used to display documentation.
func (c *ElementLib) Section(parent string, classes ...string) *gouidom.Element {
	return c.newElement(&ElementCFG{
		Parent: parent,
		Typ:    gouidom.HTMLTag.Section,
		Class:  classes,
	})
}

// Pre Adds a new <pre> tag to the dom>
func (c *ElementLib) Pre(parent string, classes ...string) *gouidom.Element {
	return c.newElement(&ElementCFG{
		Parent: parent,
		Typ:    "pre",
		Class:  classes,
	})
}

// Code Adds a new <pre> tag to the dom>
func (c *ElementLib) Code(parent string, classes ...string) *gouidom.Element {
	return c.newElement(&ElementCFG{
		Parent: parent,
		Typ:    "code",
		Class:  classes,
	})
}

// Paragraph adds a <p> tag to the dom
func (c *ElementLib) Paragraph(parent string, classes ...string) *gouidom.Element {
	return c.newElement(&ElementCFG{
		Parent: parent,
		Typ:    "p",
		Class:  classes,
	})
}

// Heading add a <h> tag to the doc with a specified size
func (c *ElementLib) Heading(parent string, size int, text string, classes ...string) *gouidom.Element {
	return c.newElement(&ElementCFG{
		Parent:             parent,
		Typ:                "h" + strconv.Itoa(size),
		Class:              classes,
		InitializationText: text,
	})
}
