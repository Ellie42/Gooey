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
	parentRect := l.GetRectAbsolute()

	if l.Prefs.Padding != nil {
		parentRect = parentRect.WithPaddingAbsolute(l.Prefs.Padding.ToDirectionalRect(Context.Resolution))
	}

	stepOffset := float32(0)

	for i := 0; i < index; i++ {
		child := l.Children[i]
		siblingRect := child.GetRectAbsolute().RelativeWithin(parentRect)
		stepOffset += siblingRect.Width
	}

	remaining := 1 - stepOffset
	remainingForChild := remaining / float32(len(l.Children)-index)

	return l.getChildRectAbsolute(index, remainingForChild, stepOffset)
}

func (l *LinearLayout) getChildRectAbsolute(index int, step float32, stepOffset float32) dimension.Rect {
	childRect := dimension.Rect{
		X: stepOffset,
		Y: 0,
		//TODO fix this when using > 2 children, it will surely overshoot
		Width:  step,
		Height: 1,
	}.RelativeToAbsolute(l.GetRectAbsolute())

	if l.Prefs.Padding != nil {
		childPadding := *l.Prefs.Padding

		if index > 0 {
			// Remove padding for all after first so only outer edge has padding
			childPadding.Left.Amount = 0
		}

		childRect = childRect.WithPaddingRelative(childPadding.ToDirectionalRect(Context.Resolution))
	}
	return childRect
}

func (l *LinearLayout) Render() {
	l.RenderStyles(l.GetRectAbsolute(), l.Prefs.StyleSettings)

	if l.Children == nil {
		return
	}

	for _, child := range l.Children {
		child.Render()
	}

	l.stepOffset = 0

	l.ShowBaseDebug()
}

func (l *LinearLayout) Init() {
	l.InitChildren(l)

	l.Initialised = true
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

	ll.AddChildWithParent(ll, widget...)
	ll.ApplyPreferences(pref)

	return ll
}
