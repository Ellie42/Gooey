package widget

import (
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
)

type ListContentWidgetConstructor func(w *WidgetListItem) Widget
type ListContentWidgetUpdater func(w Widget, index int)

type WidgetListItem struct {
	BaseWidget
	Parent            *List
	child             Widget
	widgetConstructor ListContentWidgetConstructor
	updater           ListContentWidgetUpdater
}

func (w *WidgetListItem) Init() {
	w.child = w.widgetConstructor(w)
	w.AddChildWithParent(w, w.child)
	w.child.Init()
}

func (w *WidgetListItem) Render() {
	w.updater(w.child, w.GetIndex())
	w.child.Render()
}

func newWidgetListItem(content ListContentWidgetConstructor, updater ListContentWidgetUpdater) *WidgetListItem {
	w := &WidgetListItem{
		widgetConstructor: content,
		updater:           updater,
	}

	w.Rect = dimension.Rect{0, 0, 1, 1}

	return w
}
