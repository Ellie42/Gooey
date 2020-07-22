package widget

import "git.agehadev.com/elliebelly/gooey/pkg/dimension"

type BaseWidget struct {
	Parent   WidgetParent
	Children []Widget
	Rect     dimension.Rect
}

func (p *BaseWidget) GetRectRelative() *dimension.Rect {
	return p.Rect.RelativeTo(*p.Parent.GetRectRelative())
}

func (p *BaseWidget) Init(parent WidgetParent) {
	p.baseInit(parent)
}

func (p *BaseWidget) baseInit(parent WidgetParent) {
	p.Parent = parent

	if p.Children == nil {
		return
	}

	for _, child := range p.Children {
		child.Init(p)
	}
}

func (p *BaseWidget) Render() {
	p.baseRender()
}

func (p *BaseWidget) baseRender() {
	if p.Children == nil {
		return
	}

	for _, child := range p.Children {
		child.Render()
	}
}
