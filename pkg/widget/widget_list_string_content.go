package widget

import (
	"git.agehadev.com/elliebelly/gooey/pkg/draw"
)

type StringDataProvider interface {
	Provide(index int) string
}

type StringContentProvider struct {
	DataProvider StringDataProvider
}

func (s StringContentProvider) InitListItem(w *WidgetListItem) {
}

func (s StringContentProvider) RenderListItem(w *WidgetListItem) {
	draw.Text(w.GetRectAbsolute(), s.DataProvider.Provide(w.GetIndex()), 32)
}

func NewStringListContentProvider(provider StringDataProvider) *StringContentProvider {
	return &StringContentProvider{
		DataProvider: provider,
	}
}
