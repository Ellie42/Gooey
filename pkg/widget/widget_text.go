package widget

import (
	"git.agehadev.com/elliebelly/gooey/pkg/draw"
	"git.agehadev.com/elliebelly/gooey/pkg/widget/settings"
)

type Text struct {
	BaseWidget
	Text           string
	FontSizePixels int
	dataProvider   func() string
}

func (t *Text) Render() {
	if t.dataProvider != nil {
		t.Text = t.dataProvider()
	}

	draw.Text(t.GetRectAbsolute(), t.Text, t.FontSizePixels)
}

func NewTextWidget(prefs *settings.WidgetPreferences, text string) *Text {
	t := &Text{
		FontSizePixels: 24,
		Text:           text,
	}

	t.ApplyPreferences(prefs)

	return t
}

func NewDynamicTextWidget(prefs *settings.WidgetPreferences, textProvider func() string) *Text {
	t := NewTextWidget(prefs, "")

	t.dataProvider = textProvider

	return t
}
