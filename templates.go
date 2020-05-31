package gouielement

import "github.com/goog-lukemc/gouidom"

type ArticleData struct {
	Title    string
	Subtitle string
	Content  []*ContentConfig
}

type ContentConfig struct {
	Typ string
	CFG map[string]interface{}
}

func (c *ElementLib) Readable(parent string, data *ArticleData) {

	article := c.Article(parent)

	c.Span(PathOf(article), data.Title, "heading")
	c.Span(PathOf(article), data.Subtitle, "subheading")

	section := c.Section(PathOf(article))

	c.SetContent(PathOf(section), data.Content)

}

func (c *ElementLib) SetContent(parent string, cc []*ContentConfig) {
	for _, item := range cc {
		switch item.Typ {
		case "span":
			c.Span(parent, item.CFG["text"].(string), item.CFG["classes"].([]string))
		case "img":
			c.IMG(parent, item.CFG["ea"].(map[string]string), item.CFG["classes"].([]string))
		default:
			gouidom.CLog("Not Implemented:%v", c)
		}
	}
}

func (c *ElementLib) CodeBlock(parent string, content string, classes ...string) {
	w := c.Div(parent, map[string]string{})
	pa := c.Pre(PathOf(w))
	c.Code(PathOf(pa), content)
}
