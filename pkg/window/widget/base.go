package widget

import "git.agehadev.com/elliebelly/gooey/pkg/dimension"

type BaseWidget struct {
	Parent   WidgetParent
	Index    int
	Children []Widget
	Rect     dimension.Rect
}

func (b *BaseWidget) GetRectRelative() *dimension.Rect {
	return b.Rect.RelativeTo(*b.Parent.GetChildRectRelative(b.Index))
}

func (b *BaseWidget) GetChildRectRelative(index int) *dimension.Rect {
	return b.GetRectRelative()
}

func (b *BaseWidget) SetIndex(index int) {
	b.Index = index
}

func (b *BaseWidget) Init(parent WidgetParent) {
	b.baseInit(parent)
}

func (b *BaseWidget) baseInit(parent WidgetParent) {
	b.Parent = parent

	if b.Children == nil {
		return
	}

	for _, child := range b.Children {
		child.Init(b)
	}
}

func (b *BaseWidget) Render() {
	b.baseRender()
}

func (b *BaseWidget) baseRender() {
	if b.Children == nil {
		return
	}

	for _, child := range b.Children {
		child.Render()
	}
}
