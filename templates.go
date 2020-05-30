package gouielement

import "github.com/goog-lukemc/gouidom"

type ReadingData struct {
	Title    string
	Subtitle string
	Content  []*gouidom.Element
}

func (c *ElementLib) Readable(parent string, data *ReadingData) {
	wrapper := c.WrapperDiv(parent)

	article := c.Article(PathOf(wrapper))

	c.Heading(PathOf(article), 1, data.Title)

	c.Heading(PathOf(article), 3, data.Subtitle)

	section := c.Section(PathOf(article))

	for _, item := range data.Content {
		section.AppendChild(item)
	}

}

func (c *ElementLib) CodeBlock(parent string, content string, classes ...string) {
	w := c.Div(parent, map[string]string{})
	pa := c.Pre(PathOf(w))
	c.Code(PathOf(pa), content)
}

// func (c *ComponentLib) NewPreCode(parent string, text string, classes ...string) *gouidom.Element {
// 	pre := c.NewPre(parent)
// 	code := c.NewCode(PathOf(pre))

// 	spl := strings.Split(text, "\n")
// 	for _, line := range spl {
// 		c.NewSpan(PathOf(code), line, "code-line")
// 	}

// 	return pre
// }

// func (c *ComponentLib) NewParagraph(parent string, text string, classes ...string) *gouidom.Element {
// 	return c.newComponent(&ComponentCFG{
// 		Parent:             parent,
// 		Typ:                "p",
// 		Class:              classes,
// 		InitializationText: text,
// 	})
// }

// func (c *ComponentLib) NewHeading(parent string, text string, size string, classes ...string) *gouidom.Element {
// 	return c.newComponent(&ComponentCFG{
// 		Parent:             parent,
// 		Typ:                "h" + size,
// 		Class:              classes,
// 		InitializationText: text,
// 	})
// }

// // NewButtonGroup Add a new set of buttons to the UI.
// func (c *ComponentLib) NewButtonGroup(parent string, header string, text ...string) *gouidom.Element {
// 	wrapper := c.NewDivWrapper(parent)

// 	// Add a standard header
// 	c.NewSpan(PathOf(wrapper), header, "bg-header")
// 	for _, bf := range text {
// 		c.NewButton(PathOf(wrapper), bf)
// 	}
// 	return wrapper
// }

// // NewSpan create a new span element in the dom with the provided text.
// func (c *ComponentLib) NewSpan(parent string, text string, classes ...string) *gouidom.Element {
// 	return c.newComponent(&ComponentCFG{
// 		Parent:             parent,
// 		Typ:                gouidom.HTMLTag.Span,
// 		InitializationText: text,
// 		Class:              classes,
// 	})
// }

// func (c *ComponentLib) NewDivWrapper(parent string, classes ...string) *gouidom.Element {
// 	return c.newComponent(&ComponentCFG{
// 		Parent: parent,
// 		Typ:    gouidom.HTMLTag.Div,
// 		Class:  classes,
// 	})
// }

// func (c *ComponentLib) NewDiv(parent string, ca map[string]string, classes ...string) *gouidom.Element {
// 	return c.newComponent(&ComponentCFG{
// 		Parent:           parent,
// 		Typ:              gouidom.HTMLTag.Div,
// 		Class:            classes,
// 		CustomAttributes: ca,
// 	})
// }

// func (c *ComponentLib) NewGrid(parent string, num int, classes ...string) []*gouidom.Element {
// 	// Wrapper for the grid
// 	wrapper := c.NewDivWrapper(parent, classes...)
// 	result := []*gouidom.Element{}
// 	for i := 0; i < num; i++ {
// 		result = append(result, c.NewDiv(PathOf(wrapper), map[string]string{}))
// 	}
// 	return result
// }
