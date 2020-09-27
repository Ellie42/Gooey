package widget

import (
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
	"git.agehadev.com/elliebelly/gooey/pkg/draw"
	"git.agehadev.com/elliebelly/gooey/pkg/widget/behaviour"
	"git.agehadev.com/elliebelly/gooey/pkg/widget/settings"
	"math"
)

type ListContentProvider interface {
	InitListItem(listItem *WidgetListItem)
	RenderListItem(listItem *WidgetListItem)
}

type List struct {
	BaseWidget

	Children []*WidgetListItem

	ContentProvider ListContentWidgetConstructor

	initialisedMap map[int]bool
	lastMaxRows    int
	Updater        ListContentWidgetUpdater
}

func (l *List) RecalculateChildRects() {
	if l.Children == nil {
		return
	}

	if len(l.childRects) < len(l.Children) {
		l.childRects = make([]dimension.Rect, len(l.Children))
	}

	for i, _ := range l.Children {
		rect := l.GetRectAbsolute()

		if l.Prefs.Padding != nil {
			rect = rect.WithPaddingAbsolute(l.Prefs.Padding.ToDirectionalRect(Context.Resolution))
		}

		resolution := Context.Resolution
		maxHeightPixels := 32

		rectHeightPx := int(rect.Height * float32(resolution.Height))
		relativeRowHeight := float32(maxHeightPixels) / float32(rectHeightPx)

		maxRows := float32(math.Ceil(float64(resolution.Height) / float64(maxHeightPixels)))
		heightStep := relativeRowHeight

		diffHeight := 1.0 / relativeRowHeight
		frac := 1 - (diffHeight - float32(int(diffHeight)))

		if (1 - 1e-6) < frac {
			frac = 0
		}

		rowRect := dimension.Rect{
			rect.X, -(frac * relativeRowHeight) + rect.Y + heightStep*((maxRows-1)-float32(i)), rect.Width, heightStep,
		}

		l.childRects[i] = rowRect
	}
	for _, c := range l.Children {
		c.RecalculateChildRects()
	}
}

//func (l *List) GetChildRectAbsolute(index int) dimension.Rect {
//	rect := l.GetRectAbsolute()
//	resolution := Context.Resolution
//	maxHeightPixels := 32
//	maxRows := float32(math.Ceil(float64(resolution.Height) / float64(maxHeightPixels)))
//	heightStep := (1 / maxRows) * rect.Height
//
//	rowRect := dimension.Rect{
//		rect.X, rect.Y + heightStep*((maxRows-1)-float32(index)), rect.Width, heightStep,
//	}
//
//	return rowRect
//}

func NewListStringContentWidget(w *WidgetListItem) Widget {
	t := NewTextWidget(nil, "")

	t.Text = "I have nothing to say"

	return t
}

func (l *List) Render() {
	resolution := Context.Resolution
	maxHeightPixels := 32
	maxRows := float32(math.Ceil(float64(resolution.Height) / float64(maxHeightPixels)))

	if int(maxRows) > len(l.Children) {
		l.Children = append(l.Children, make([]*WidgetListItem, int(maxRows)-len(l.Children))...)
	}

	newChildren := false

	for i := 0; i < int(maxRows); i++ {
		if l.Children[i] == nil {
			listItem := newWidgetListItem(l.ContentProvider, l.Updater)

			listItem.SetParent(l)
			listItem.SetIndex(i)

			listItem.Init()

			l.Children[i] = listItem
			newChildren = true
		}
	}

	if newChildren {
		l.RecalculateChildRects()
	}

	for i, child := range l.Children {
		rowRect := l.GetChildRectAbsolute(i)

		colours := []draw.RGBA{
			draw.NewRGBAFromHex("0e0e10"),
			draw.NewRGBAFromHex("16161a"),
		}

		draw.SquareFilled(rowRect, colours[i%2])

		child.Render()
	}
}

func NewList(provider ListContentWidgetConstructor, updater ListContentWidgetUpdater, prefs *settings.WidgetPreferences) *List {
	list := &List{
		ContentProvider: provider,
		Updater:         updater,
		initialisedMap:  make(map[int]bool),
	}

	if prefs == nil {
		prefs = &settings.WidgetPreferences{}
	}

	prefs.Behaviours.Scrollable = &behaviour.Scrollable{}

	list.Rect.Width = 1
	list.Rect.Height = 1

	list.ApplyPreferences(prefs)

	return list
}
