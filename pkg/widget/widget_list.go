package widget

import (
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
	"git.agehadev.com/elliebelly/gooey/pkg/draw"
	"git.agehadev.com/elliebelly/gooey/pkg/widget/settings"
	"math"
)

type ListContentProvider interface {
	InitListItem(listItem *WidgetListItem)
	RenderListItem(listItem *WidgetListItem)
}

type List struct {
	BaseWidget

	ContentProvider ListContentProvider

	initialisedMap map[int]bool
	lastMaxRows    int
}

func (l *List) GetChildRectAbsolute(index int) dimension.Rect {
	rect := l.GetRectAbsolute()
	resolution := Context.Resolution
	maxHeightPixels := 32
	maxRows := float32(math.Ceil(float64(resolution.Height) / float64(maxHeightPixels)))
	heightStep := (1 / maxRows) * rect.Height

	rowRect := dimension.Rect{
		rect.X, rect.Y + heightStep*((maxRows-1)-float32(index)), rect.Width, heightStep,
	}

	return rowRect
}

func (l *List) Init() {
	l.InitChildren(l)
}

func (l *List) Render() {
	resolution := Context.Resolution
	maxHeightPixels := 32
	maxRows := float32(math.Ceil(float64(resolution.Height) / float64(maxHeightPixels)))

	if int(maxRows) > len(l.Children) {
		l.Children = append(l.Children, make([]Widget, int(maxRows)-len(l.Children))...)
	}

	for i := 0; i < int(maxRows); i++ {
		rowRect := l.GetChildRectAbsolute(i)

		colours := []draw.RGBA{
			draw.NewRGBAFromHex("A14643"),
			draw.NewRGBAFromHex("A1514E"),
		}

		draw.SquareFilled(rowRect, colours[i%2])

		if l.Children[i] == nil {
			listItem := NewWidgetListItem(l.ContentProvider)

			listItem.SetParent(l)
			listItem.SetIndex(i)

			listItem.Init()

			l.Children[i] = listItem
		}

		l.Children[i].Render()
	}
}

func NewList(provider ListContentProvider, prefs *settings.WidgetPreferences) *List {
	list := &List{
		ContentProvider: provider,
		initialisedMap:  make(map[int]bool),
	}

	list.Rect.Width = 1
	list.Rect.Height = 1

	list.ApplyPreferences(prefs)

	return list
}
