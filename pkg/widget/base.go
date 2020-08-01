package widget

import (
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
	"git.agehadev.com/elliebelly/gooey/pkg/draw"
	"git.agehadev.com/elliebelly/gooey/pkg/widget/settings"
	"git.agehadev.com/elliebelly/gooey/pkg/widget/styles"
	"math/rand"
)

type BaseWidget struct {
	Prefs settings.WidgetPreferences
	styles.Styles

	Parent      Widget
	Index       int
	Children    []Widget
	debugColour *draw.RGBA
	Rect        dimension.Rect
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

	bindRectDimensionToSize(&rect, b.Prefs.DimensionBounds, parentRect, Width)
	bindRectDimensionToSize(&rect, b.Prefs.DimensionBounds, parentRect, Height)

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
		rect = rect.RelativeTo(parentRect)
	}

	return rect
}

type Dimension int

const (
	Width Dimension = iota
	Height
)

func bindRectDimensionToSize(rect *dimension.Rect, bounds *dimension.Dimensions, parentRect dimension.Rect, dim Dimension) {
	var pixelDimensionValue, parentRectDimensionValue float32
	var windowResolutionDimensionValue int
	var size *dimension.Size
	var value *float32

	if bounds == nil {
		return
	}

	if rect == nil {
		panic("bind rect dimension failed, rect cannot be nil")
	}

	pixelPos := rect.RelativeTo(parentRect).MultipliedByDimension(Context.Resolution)

	switch dim {
	case Width:
		pixelDimensionValue = pixelPos.Width
		parentRectDimensionValue = parentRect.Width
		windowResolutionDimensionValue = Context.Resolution.Width
		size = bounds.Width
		value = &rect.Width
	case Height:
		pixelDimensionValue = pixelPos.Height
		parentRectDimensionValue = parentRect.Height
		windowResolutionDimensionValue = Context.Resolution.Height
		size = bounds.Height
		value = &rect.Height
	}

	if size == nil || value == nil {
		return
	}

	switch size.Unit {
	case dimension.SizeUnitPixels:
		if pixelDimensionValue > size.Amount {
			*value = size.Amount / float32(windowResolutionDimensionValue) / parentRectDimensionValue
		}
	case dimension.SizeUnitRatio:
		if *value > size.Amount {
			*value = size.Amount
		}
	}
}

func (b *BaseWidget) GetChildRectAbsolute(index int) dimension.Rect {
	rect := b.GetRectAbsolute()

	if b.Prefs.Padding != nil {
		rect = rect.WithPadding(*b.Prefs.Padding)
	}

	return rect
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
	if p != nil {
		b.Prefs = *p

		if p.Rect != nil {
			b.Rect = *p.Rect
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

func (b *BaseWidget) AddChild(widget ...Widget) {
	b.Children = append(b.Children, widget...)
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
		draw.SquareEdge(b.GetRectAbsolute().Shrink(0.01), *b.debugColour)
	}
}
