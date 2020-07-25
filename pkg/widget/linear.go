package widget

import (
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
	"git.agehadev.com/elliebelly/gooey/pkg/widget/settings"
)

type LinearLayout struct {
	BaseWidget
	Initialised bool
}

func (l *LinearLayout) GetChildRectAbsolute(index int) dimension.Rect {
	step := float32(1) / float32(len(l.Children))

	childRect := dimension.Rect{
		X:      float32(index) * step,
		Y:      0,
		Width:  step,
		Height: 1,
	}.RelativeTo(l.GetRectAbsolute())

	return childRect
}

func (l *LinearLayout) Init() {
	l.InitChildren(l)

	l.Initialised = true
}

func (ll *LinearLayout) AddChild(widget ...Widget) {
	for _, wid := range widget {
		wid.SetIndex(len(ll.Children))
		ll.Children = append(ll.Children, wid)
	}
}

func NewLinearLayout(pref *settings.WidgetPreferences, widget ...Widget) *LinearLayout {
	ll := &LinearLayout{}

	ll.Rect = &dimension.Rect{
		X:      0,
		Y:      0,
		Width:  1,
		Height: 1,
	}

	ll.AddChild(widget...)
	ll.ApplyPreferences(pref)

	return ll
}

func addChildren(widget []Widget, ll *LinearLayout) {

}
