package widget

import "git.agehadev.com/elliebelly/gooey/lib/dimension"

type Widget interface {
	SetIndex(index int)
	SetParent(parent Widget)
	Init()
	AddChild(widget ...Widget)
	GetRectAbsolute() dimension.Rect
	GetChildRectAbsolute(index int) dimension.Rect
	Render()
}
