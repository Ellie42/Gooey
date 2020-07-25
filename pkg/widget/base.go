package widget

import (
	"git.agehadev.com/elliebelly/gooey/internal/renderer/draw"
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
	"git.agehadev.com/elliebelly/gooey/pkg/widget/settings"
)

type BaseWidget struct {
	settings.WidgetPreferences

	Parent   Widget
	Index    int
	Children []Widget
}

func (b *BaseWidget) GetRectAbsolute() dimension.Rect {
	var rect dimension.Rect

	if b.Rect == nil {
		rect = dimension.Rect{
			0, 0, 1, 1,
		}
	} else {
		rect = dimension.Rect{
			X:      b.Rect.X,
			Y:      b.Rect.Y,
			Width:  b.Rect.Width,
			Height: b.Rect.Height,
		}
	}

	switch b.FixedRatioAxis {
	case settings.FixedX:
		rect.Width = rect.Height * b.FixedRatio * (float32(Context.Resolution.Height) / float32(Context.Resolution.Width))
	case settings.FixedY:
		rect.Height = rect.Width * b.FixedRatio * (float32(Context.Resolution.Width) / float32(Context.Resolution.Height))
	}

	switch b.AlignmentVertical {
	case settings.VerticalMiddle:
		rect.Y = (1 - rect.Height) / 2
	case settings.VerticalTop:
		rect.Y = 1 - rect.Height
	case settings.VerticalBottom:
		rect.Y = 0
	}

	switch b.AlignmentHorizontal {
	case settings.HorizontalMiddle:
		rect.X = (1 - rect.Width) / 2
	case settings.HorizontalLeft:
		rect.X = 0
	case settings.HorizontalRight:
		rect.X = 1 - rect.Width
	}

	if b.Padding != nil {
		rect = rect.WithPadding(*b.Padding)
	}

	if b.Parent != nil {
		rect = rect.RelativeTo(b.Parent.GetChildRectAbsolute(b.Index))
	}

	return rect
}

func (b *BaseWidget) GetChildRectAbsolute(index int) dimension.Rect {
	return b.GetRectAbsolute()
}

func (b *BaseWidget) SetIndex(index int) {
	b.Index = index
}

func (b *BaseWidget) SetParent(parent Widget) {
	b.Parent = parent
}

func (b *BaseWidget) Init() {
	b.InitChildren(b)
}

func (b *BaseWidget) InitChildren(parent Widget) {
	if b.Children == nil {
		return
	}

	for _, child := range b.Children {
		child.SetParent(parent)
		child.Init()
	}
}

func (b *BaseWidget) ApplyPreferences(p *settings.WidgetPreferences) {
	if p == nil {
		return
	}

	if p.Rect != nil {
		b.Rect = p.Rect
	}

	if p.Padding != nil {
		b.Padding = p.Padding
	}

	b.FixedRatio = p.FixedRatio
	b.FixedRatioAxis = p.FixedRatioAxis
	b.AlignmentHorizontal = p.AlignmentHorizontal
	b.AlignmentVertical = p.AlignmentVertical
}

func (b *BaseWidget) Render() {
	b.RenderChildren()
	b.ShowBaseDebug()
}

func (b *BaseWidget) RenderChildren() {
	if b.Children == nil {
		return
	}

	for _, child := range b.Children {
		child.Render()
	}
}

func (b *BaseWidget) AddChild(widget ...Widget) {
	b.Children = append(b.Children, widget...)
}

func (b *BaseWidget) ShowBaseDebug() {
	if Context.ShowDebug {
		draw.Square(b.GetRectAbsolute())
	}
}
