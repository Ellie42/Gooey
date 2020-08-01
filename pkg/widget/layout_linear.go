package widget

import (
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
	"git.agehadev.com/elliebelly/gooey/pkg/widget/settings"
)

type LinearLayout struct {
	BaseWidget
	Initialised   bool
	FitToChildren bool
	stepOffset    float32
}

func (l *LinearLayout) GetChildRectAbsolute(index int) dimension.Rect {
	step := float32(1) / float32(len(l.Children))

	childRect := dimension.Rect{
		X: float32(index)*step + l.stepOffset,
		Y: 0,
		//TODO fix this when using > 2 children, it will surely overshoot
		Width:  step - l.stepOffset,
		Height: 1,
	}.RelativeTo(l.GetRectAbsolute())

	if l.Prefs.Padding != nil {
		childPadding := *l.Prefs.Padding

		if index > 0 {
			childPadding.Left = 0
		}

		childRect = childRect.WithPadding(childPadding)
	}

	return childRect
}

func (l *LinearLayout) Render() {
	l.RenderStyles(l.GetRectAbsolute(), l.Prefs.StyleSettings)

	if l.Children == nil {
		return
	}

	for i, child := range l.Children {
		child.Render()

		rect := l.GetChildRectAbsolute(i)

		if l.FitToChildren {
			childAbsRect := child.GetRectAbsolute()

			if childAbsRect.Width < rect.Width {
				l.stepOffset -= (rect.Width - childAbsRect.Width)
			}
		}
	}

	l.stepOffset = 0

	l.ShowBaseDebug()
}

func (l *LinearLayout) Init() {
	l.InitChildren(l)

	l.Initialised = true
}

func (ll *LinearLayout) AddChild(widget ...Widget) {
	for _, wid := range widget {
		wid.SetIndex(len(ll.Children))
		wid.SetParent(ll)
		ll.Children = append(ll.Children, wid)
	}
}

func NewLinearLayout(pref *settings.WidgetPreferences, widget ...Widget) *LinearLayout {
	ll := &LinearLayout{}

	ll.Prefs.Rect = &dimension.Rect{
		X:      0,
		Y:      0,
		Width:  1,
		Height: 1,
	}

	ll.FitToChildren = true

	ll.AddChild(widget...)
	ll.ApplyPreferences(pref)

	return ll
}

