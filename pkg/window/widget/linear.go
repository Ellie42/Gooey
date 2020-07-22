package widget

import (
	"git.agehadev.com/elliebelly/gooey/pkg/dimension"
)

type WidgetParent interface {
	GetRectRelative() *dimension.Rect
}

type LinearLayout struct {
	BaseWidget
	Initialised bool
}

func (l *LinearLayout) Add(widget ...Widget) {
	l.Children = append(l.Children, widget...)
}

func (l *LinearLayout) GetRectRelative() *dimension.Rect {
	return l.Parent.GetRectRelative()
}

func (l *LinearLayout) Render() {
	for _, w := range l.Children {
		w.Render()
	}
}

func (l *LinearLayout) Init(parent WidgetParent) {
	l.Parent = parent

	for _, w := range l.Children {
		w.Init(l)
	}

	l.Initialised = true
}

func NewLinearLayout() *LinearLayout {
	ll := &LinearLayout{}

	ll.Children = make([]Widget, 0)

	return ll
}
