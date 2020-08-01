package styles

import (
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
	"git.agehadev.com/elliebelly/gooey/pkg/draw"
)

type Styles struct {
}

func (s *Styles) RenderStyles(rect dimension.Rect, settings *StyleSettings) {
	if settings == nil {
		return
	}

	if settings.BackgroundColour != nil {
		draw.SquareFilled(rect, *settings.BackgroundColour)
	}
}
