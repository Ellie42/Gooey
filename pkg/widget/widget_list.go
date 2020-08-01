package widget

import (
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
	"git.agehadev.com/elliebelly/gooey/pkg/draw"
	"git.agehadev.com/elliebelly/gooey/pkg/widget/settings"
	"math"
)

type List struct {
	BaseWidget
}

func (l *List) Render() {
	rect := l.GetRectAbsolute()
	resolution := Context.Resolution
	maxHeightPixels := 30
	//maxHeightPixels := 300
	maxRows := float32(math.Ceil(float64(resolution.Height) / float64(maxHeightPixels)))
	heightStep := (1 / maxRows) * rect.Height

	colours := []draw.RGBA{
		draw.NewRGBAFromHex("872321"),
		draw.NewRGBAFromHex("8A2421"),
	}

	for i := 0; i < int(maxRows); i++ {
		rowRect := dimension.Rect{
			rect.X, rect.Y + heightStep*float32(i), rect.Width, heightStep,
		}

		draw.SquareFilled(rowRect, colours[i%2])

		draw.Text(rowRect, "There is text here, there and everywhere", 32)
	}
}

func NewList(prefs *settings.WidgetPreferences) *List {
	list := &List{}

	list.Rect.Width = 1
	list.Rect.Height = 1

	list.ApplyPreferences(prefs)

	return list
}
