package gouicomponent

import (
	"path"

	"github.com/goog-lukemc/gouidom"
)

type ComponentLib struct {
	v *gouidom.VDOM
}

type ComponentCFG struct {
	ID                 string
	Parent             string
	Typ                string
	InitializationText string
	Class              []string
}

func NewComponentLib(v *gouidom.VDOM) *ComponentLib {
	return &ComponentLib{v: v}
}

// NewSpan create a new span element in the dom with the provided text.
func (c *ComponentLib) NewSpan(parent string, text string, classes ...string) *gouidom.Element {
	return c.newConponent(&ComponetCFG{
		Parent:             parent,
		Typ:                gouidom.HTMLTag.Span,
		InitializationText: text,
		Class:              classes,
	})
}

func (c *ComponentLib) NewWrapperDiv(parent string) *gouidom.Element {
	return c.newComponent(&ComponentCFG{
		Parent: parent,
		Typ:    gouidom.HTMLTag.Div,
		Class:  []string{"grid", "container"},
	})
}

func (c *ComponentLib) newConponent(cfg *ComponetCFG) *gouidom.Element {
	ele, err := gouidom.NewElement(cfg.ID, cfg.Parent, cfg.Typ, cfg.InitializationText, cfg.Class...)
	if err != nil {
		gouidom.CLog("%s", err.Error())
	}
	c.v.AddElement(ele)
	return ele
}

func NewWrapper(parent string, v *gouidom.VDOM) *gouidom.Element {
	wrapper, err := gouidom.NewElement("", parent, gouidom.HTMLTag.Div, "", "grid", "container")
	if err != nil {
		gouidom.CLog("%s", err.Error())
	}
	v.AddElement(wrapper)
	return wrapper
}

func Grid(parent string, num int, v *gouidom.VDOM, class ...string) {
	// Create the div wrapper for the layout
	wrapper, err := gouidom.NewElement("", parent, gouidom.HTMLTag.Div, "", "grid", "container")
	if err != nil {
		gouidom.CLog("%s", err.Error())
	}
	v.AddElement(wrapper)

	for i := 0; i < num; i++ {
		content, err := gouidom.NewElement("", path.Join(wrapper.Parent, wrapper.ID), gouidom.HTMLTag.Div, "", class...)
		if err != nil {
			gouidom.CLog("%s", err.Error())
		}
		v.AddElement(content)
	}

}
