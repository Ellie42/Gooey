package widget

import (
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
	"git.agehadev.com/elliebelly/gooey/pkg/draw"
	"git.agehadev.com/elliebelly/gooey/pkg/widget/behaviour"
	"git.agehadev.com/elliebelly/gooey/pkg/widget/settings"
	"git.agehadev.com/elliebelly/gooey/pkg/widget/styles"
	"math/rand"
)

type BaseWidget struct {
	Prefs      settings.WidgetPreferences
	Behaviours behaviour.BehaviourSet
	styles.Styles

	Parent      Widget
	Index       int
	Children    []Widget
	debugColour *draw.RGBA
	Rect        dimension.Rect
	childRects  []dimension.Rect
}

func (b *BaseWidget) GetPrefs() *settings.WidgetPreferences {
	return &b.Prefs
}

func (b *BaseWidget) RecalculateChildRects() {
	if b.Children == nil {
		return
	}

	if len(b.childRects) < len(b.Children) {
		b.childRects = make([]dimension.Rect, len(b.Children))
	}

	for i, _ := range b.Children {
		rect := b.GetRectAbsolute()

		if b.Prefs.Padding != nil {
			rect = rect.WithPaddingAbsolute(b.Prefs.Padding.ToDirectionalRect(Context.Resolution))
		}

		b.childRects[i] = rect
	}
	for _, c := range b.Children {
		c.RecalculateChildRects()
	}
}

func (b *BaseWidget) GetRectAbsolute() dimension.Rect {
	var rect dimension.Rect

	if b.Prefs.Rect == nil {
		rect = dimension.Rect{
			0, 0, 1, 1,
		}
	} else {
		rect = dimension.Rect{
			X:      b.Prefs.Rect.X,
			Y:      b.Prefs.Rect.Y,
			Width:  b.Prefs.Rect.Width,
			Height: b.Prefs.Rect.Height,
		}
	}

	var parentRect dimension.Rect

	if b.Parent != nil {
		parentRect = b.Parent.GetChildRectAbsolute(b.Index)
	}

	if b.Prefs.DimensionBounds != nil {
		b.Prefs.DimensionBounds.Bind(&rect, parentRect, dimension.Width, Context.Resolution)
		b.Prefs.DimensionBounds.Bind(&rect, parentRect, dimension.Height, Context.Resolution)
	}

	switch b.Prefs.FixedRatioAxis {
	case settings.FixedX:
		rect.Width = rect.Height * b.Prefs.FixedRatio *
			(float32(Context.Resolution.Height) / float32(Context.Resolution.Width)) *
			parentRect.GetRatioX()
	case settings.FixedY:
		rect.Height = rect.Width * b.Prefs.FixedRatio *
			(float32(Context.Resolution.Width) / float32(Context.Resolution.Height)) *
			parentRect.GetRatioY()
	}

	switch b.Prefs.AlignmentVertical {
	case settings.VerticalMiddle:
		rect.Y = (1 - rect.Height) / 2
	case settings.VerticalTop:
		rect.Y = 1 - rect.Height
	case settings.VerticalBottom:
		rect.Y = 0
	}

	switch b.Prefs.AlignmentHorizontal {
	case settings.HorizontalMiddle:
		rect.X = (1 - rect.Width) / 2
	case settings.HorizontalLeft:
		rect.X = 0
	case settings.HorizontalRight:
		rect.X = 1 - rect.Width
	}

	if b.Parent != nil {
		rect = rect.RelativeToAbsolute(parentRect)
	}

	return rect
}

func (b *BaseWidget) GetChildRectAbsolute(index int) dimension.Rect {
	return b.childRects[index]
}

func (b *BaseWidget) SetIndex(index int) {
	b.Index = index
}

func (b *BaseWidget) SetParent(parent Widget) {
	b.Parent = parent
}

func (b *BaseWidget) Init() {
	b.Behaviours.Init(&Context.EventManager, b.GetRectAbsolute)
	b.InitChildren(b)
}

func (b *BaseWidget) InitChildren(parent Widget) {
	if b.Children == nil {
		return
	}

	for _, child := range b.Children {
		child.Init()
	}
}

func (b *BaseWidget) ApplyPreferences(p *settings.WidgetPreferences) {
	if p != nil {
		b.Prefs = *p

		if p.Rect != nil {
			b.Rect = *p.Rect
		} else {
			p.Rect = &dimension.Rect{0, 0, 1, 1}
		}

		if p.Behaviours != nil {
			b.Behaviours = *b.Prefs.Behaviours
		}
	}
}

func (b *BaseWidget) Render() {
	b.RenderStyles(
		b.GetRectAbsolute(),
		b.Prefs.StyleSettings,
	)
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

func (b *BaseWidget) GetIndex() int {
	return b.Index
}

func (b *BaseWidget) AddChildWithParent(parent Widget, children ...Widget) {
	for _, child := range children {
		child.SetIndex(len(b.Children))
		child.SetParent(parent)
		b.Children = append(b.Children, child)
		b.childRects = append(b.childRects, dimension.Rect{})
	}
}

func (b *BaseWidget) AddChild(widget ...Widget) {
	b.AddChildWithParent(b, widget...)
}

func (b *BaseWidget) ShowBaseDebug() {
	if b.debugColour == nil {
		b.debugColour = &draw.RGBA{
			R: rand.Float32(),
			G: rand.Float32(),
			B: rand.Float32(),
			A: 1,
		}
	}

	if Context.ShowDebug {
		draw.SquareEdge(b.GetRectAbsolute(), *b.debugColour)
	}
}
