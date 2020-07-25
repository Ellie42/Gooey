package settings

import "git.agehadev.com/elliebelly/gooey/lib/dimension"

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
	Rect    *dimension.Rect
	Padding *dimension.DirectionalRect

	FixedRatioAxis FixedAspectAxis
	FixedRatio     float32

	AlignmentVertical
	AlignmentHorizontal
}
