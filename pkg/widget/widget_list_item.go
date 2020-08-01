package widget

import (
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
)

type WidgetListItem struct {
	BaseWidget
	Parent          *List
	ContentProvider ListContentProvider
}

func (w *WidgetListItem) Init() {
	w.ContentProvider.InitListItem(w)
}

func (w *WidgetListItem) Render() {
	w.ContentProvider.RenderListItem(w)
}

func NewWidgetListItem(content ListContentProvider) *WidgetListItem {
	w := &WidgetListItem{
		ContentProvider: content,
	}

	w.Rect = dimension.Rect{0, 0, 1, 1}

	return w
}
