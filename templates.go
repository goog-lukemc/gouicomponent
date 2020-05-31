package gouielement

import (
	"github.com/goog-lukemc/gouidom"
)

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

	gas := func(name string, cfg map[string]interface{}) []string {
		return cfg[name].([]string)
	}

	gs := func(name string, cfg map[string]interface{}) string {
		return cfg[name].(string)
	}

	gmss := func(name string, cfg map[string]interface{}) map[string]string {
		return cfg[name].(map[string]string)
	}

	for _, item := range cc {
		switch item.Typ {
		case "span":
			c.Span(parent, gs("text", item.CFG))
		case "img":
			c.IMG(parent, gmss("ea", item.CFG), gas("classes", item.CFG)...)
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
