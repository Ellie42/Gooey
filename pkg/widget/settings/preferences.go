package settings

import (
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
	"git.agehadev.com/elliebelly/gooey/pkg/widget/styles"
)

type FixedAspectAxis int

const (
	None FixedAspectAxis = iota
	FixedX
	FixedY
)

type AlignmentHorizontal int

const (
	HorizontalNone AlignmentHorizontal = iota
	HorizontalLeft
	HorizontalRight
	HorizontalMiddle
)

type AlignmentVertical int

const (
	VerticalNone AlignmentVertical = iota
	VerticalTop
	VerticalBottom
	VerticalMiddle
)

type WidgetPreferences struct {
	Rect            *dimension.Rect
	Padding         *dimension.DirectionalRectSized
	DimensionBounds *dimension.Dimensions
	StyleSettings   *styles.StyleSettings

	FixedRatioAxis FixedAspectAxis
	FixedRatio     float32

	AlignmentVertical
	AlignmentHorizontal
}
