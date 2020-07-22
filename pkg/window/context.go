package window

import "git.agehadev.com/elliebelly/gooey/pkg/dimension"

type WindowContext struct {
	Resolution dimension.SizeInt
	Rect       dimension.Rect
}

var Context *WindowContext
