package gouicomponent

import (
	"path"
	"strings"

	"github.com/goog-lukemc/gouidom"
)

func (c *ComponentLib) NewPreCode(parent string, text string, classes ...string) *gouidom.Element {
	pre := c.NewPre(parent)
	code := c.NewCode(PathOf(pre))

	spl := strings.Split(text, "\n")
	for _, line := range spl {
		c.NewSpan(PathOf(code), line, "code-line")
	}

	return pre
}

func (c *ComponentLib) NewCode(parent string, classes ...string) *gouidom.Element {
	return c.newComponent(&ComponentCFG{
		Parent: parent,
		Typ:    "code",
		Class:  classes,
	})
}

func (c *ComponentLib) NewPre(parent string, classes ...string) *gouidom.Element {
	return c.newComponent(&ComponentCFG{
		Parent: parent,
		Typ:    "pre",
		Class:  classes,
	})
}

func (c *ComponentLib) NewParagraph(parent string, text string, classes ...string) *gouidom.Element {
	return c.newComponent(&ComponentCFG{
		Parent:             parent,
		Typ:                "p",
		Class:              classes,
		InitializationText: text,
	})
}

func (c *ComponentLib) NewHeading(parent string, text string, size string, classes ...string) *gouidom.Element {
	return c.newComponent(&ComponentCFG{
		Parent:             parent,
		Typ:                "h" + size,
		Class:              classes,
		InitializationText: text,
	})
}

// NewSection Adds and article tag usually used to display documentation.
func (c *ComponentLib) NewSection(parent string, classes ...string) *gouidom.Element {
	return c.newComponent(&ComponentCFG{
		Parent: parent,
		Typ:    gouidom.HTMLTag.Section,
		Class:  classes,
	})
}

// NewArticle Adds and article tag usually used to display documentation.
func (c *ComponentLib) NewArticle(parent string, classes ...string) *gouidom.Element {
	return c.newComponent(&ComponentCFG{
		Parent: parent,
		Typ:    gouidom.HTMLTag.Article,
		Class:  classes,
	})
}

// NewButtonGroup Add a new set of buttons to the UI.
func (c *ComponentLib) NewButtonGroup(parent string, header string, text ...string) *gouidom.Element {
	wrapper := c.NewDivWrapper(parent)

	// Add a standard header
	c.NewSpan(PathOf(wrapper), header, "bg-header")
	for _, bf := range text {
		c.NewButton(PathOf(wrapper), bf)
	}
	return wrapper
}

// NewButton Adds a new button to a document
func (c *ComponentLib) NewButton(parent string, text string) *gouidom.Element {
	wrapper := c.NewDivWrapper(parent)
	return c.newComponent(&ComponentCFG{
		Parent:             PathOf(wrapper),
		Typ:                gouidom.HTMLTag.Button,
		InitializationText: text,
	})
}

// NewStyle injects a new style tag into the body of the document
func (c *ComponentLib) NewStyle(text string) *gouidom.Element {

	return c.newComponent(&ComponentCFG{
		Parent:             "html/body",
		Typ:                gouidom.HTMLTag.Style,
		InitializationText: text,
	})
}

// NewSpan create a new span element in the dom with the provided text.
func (c *ComponentLib) NewSpan(parent string, text string, classes ...string) *gouidom.Element {
	return c.newComponent(&ComponentCFG{
		Parent:             parent,
		Typ:                gouidom.HTMLTag.Span,
		InitializationText: text,
		Class:              classes,
	})
}

func (c *ComponentLib) NewDivWrapper(parent string, classes ...string) *gouidom.Element {
	return c.newComponent(&ComponentCFG{
		Parent: parent,
		Typ:    gouidom.HTMLTag.Div,
		Class:  classes,
	})
}

func (c *ComponentLib) NewDiv(parent string, ca map[string]string, classes ...string) *gouidom.Element {
	return c.newComponent(&ComponentCFG{
		Parent:           parent,
		Typ:              gouidom.HTMLTag.Div,
		Class:            classes,
		CustomAttributes: ca,
	})
}

func (c *ComponentLib) NewGrid(parent string, num int, classes ...string) []*gouidom.Element {
	// Wrapper for the grid
	wrapper := c.NewDivWrapper(parent, classes...)
	result := []*gouidom.Element{}
	for i := 0; i < num; i++ {
		result = append(result, c.NewDiv(PathOf(wrapper), map[string]string{}))
	}
	return result
}

func (c *ComponentLib) newComponent(cfg *ComponentCFG) *gouidom.Element {
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

func PathOf(ele *gouidom.Element) string {
	return path.Join(ele.Parent, ele.ID)
}

// func NewWrapper(parent string, v *gouidom.VDOM) *gouidom.Element {
// 	wrapper, err := gouidom.NewElement("", parent, gouidom.HTMLTag.Div, "", "grid", "container")
// 	if err != nil {
// 		gouidom.CLog("%s", err.Error())
// 	}
// 	v.AddElement(wrapper)
// 	return wrapper
// }

// func Grid(parent string, num int, v *gouidom.VDOM, class ...string) {
// 	// Create the div wrapper for the layout
// 	wrapper, err := gouidom.NewElement("", parent, gouidom.HTMLTag.Div, "", "grid", "container")
// 	if err != nil {
// 		gouidom.CLog("%s", err.Error())
// 	}
// 	v.AddElement(wrapper)

// 	for i := 0; i < num; i++ {
// 		content, err := gouidom.NewElement("", path.Join(wrapper.Parent, wrapper.ID), gouidom.HTMLTag.Div, "", class...)
// 		if err != nil {
// 			gouidom.CLog("%s", err.Error())
// 		}
// 		v.AddElement(content)
// 	}

// }
