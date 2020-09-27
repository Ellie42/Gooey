package widget

import (
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
	"git.agehadev.com/elliebelly/gooey/pkg/widget/settings"
)

type Widget interface {
	SetIndex(index int)
	GetIndex() int
	SetParent(parent Widget)
	Init()
	AddChild(widget ...Widget)
	GetRectAbsolute() dimension.Rect
	GetChildRectAbsolute(index int) dimension.Rect
	Render()
	RecalculateChildRects()
	GetPrefs() *settings.WidgetPreferences
}
