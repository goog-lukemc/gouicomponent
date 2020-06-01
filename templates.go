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
		if val, ok := cfg[name].([]string); ok {
			return val
		}
		return nil
	}

	gs := func(name string, cfg map[string]interface{}) string {
		if val, ok := cfg[name].(string); ok {
			return val
		}
		return ""
	}

	gmss := func(name string, cfg map[string]interface{}) map[string]string {
		if val, ok := cfg[name].(map[string]string); ok {
			return val
		}
		return map[string]string{}
	}

	for _, item := range cc {
		switch item.Typ {
		case "span":
			if sh := gs("secheading", item.CFG); sh != "" {
				c.Span(parent, sh, "sectionheading")
			}
			c.Span(parent, gs("text", item.CFG))
		case "img":
			c.IMG(parent, gmss("ea", item.CFG), gas("classes", item.CFG)...)
		default:
			gouidom.CLog("Not Implemented:%v", c)
		}
	}
}

// Codebock creates a prepared code block for use
func (c *ElementLib) CodeBlock(parent string, content string, classes ...string) {
	w := c.Div(parent, map[string]string{})
	pa := c.Pre(PathOf(w))
	c.Code(PathOf(pa), content)
}

// UnOrderedList create an unordered list
func (c *ElementLib) UnOrderedList(parent string, list []string) {
	w := c.Div(parent, map[string]string{})
	ul := c.Ul(PathOf(w), map[string]string{})
	for _, item := range list {
		c.Li(PathOf(ul), item)
	}
}
