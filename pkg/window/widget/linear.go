package widget

import (
	"git.agehadev.com/elliebelly/gooey/pkg/dimension"
)

type WidgetParent interface {
	GetRectRelative() *dimension.Rect
	GetChildRectRelative(index int) *dimension.Rect
}

type LinearLayout struct {
	BaseWidget
	Initialised bool
}

func (l *LinearLayout) Add(widget ...Widget) {
	l.Children = append(l.Children, widget...)
}

func (l *LinearLayout) Render() {
	for _, w := range l.Children {
		w.Render()
	}
}

func (l *LinearLayout) GetChildRectRelative(index int) *dimension.Rect {
	step := float32(1) / float32(len(l.Children))

	childRect := dimension.Rect{
		X:      float32(index) * step,
		Y:      0,
		Width:  step,
		Height: 1,
	}.RelativeTo(*l.GetRectRelative())

	return childRect
}

func (l *LinearLayout) Init(parent WidgetParent) {
	l.Parent = parent

	for _, w := range l.Children {
		w.Init(l)
	}

	l.Initialised = true
}

func NewLinearLayout(widget ...Widget) *LinearLayout {
	ll := &LinearLayout{}

	ll.Rect = dimension.Rect{
		X:      0,
		Y:      0,
		Width:  1,
		Height: 1,
	}

	ll.Children = widget

	return ll
}
