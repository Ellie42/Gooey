package widget

import (
	"git.agehadev.com/elliebelly/gooey/pkg/draw"
	"git.agehadev.com/elliebelly/gooey/pkg/widget/settings"
)

type Text struct {
	BaseWidget
	Text string
	FontSizePixels int
}

func (t *Text) Render() {
	draw.Text(t.GetRectAbsolute(), t.Text, t.FontSizePixels)
}

func NewTextWidget(prefs *settings.WidgetPreferences) *Text {
	t := &Text{
		FontSizePixels: 32,
	}

	t.ApplyPreferences(prefs)

	return t
}
