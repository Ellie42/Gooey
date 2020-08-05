package widget

import (
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
	"git.agehadev.com/elliebelly/gooey/pkg/widget/settings"
)

type FreeLayout struct {
	BaseWidget
}

func NewFreeLayout(pref *settings.WidgetPreferences, widget ...Widget) *FreeLayout {
	ll := &FreeLayout{}

	ll.Prefs.Rect = &dimension.Rect{
		X:      0,
		Y:      0,
		Width:  1,
		Height: 1,
	}

	ll.AddChildWithParent(ll, widget...)

	ll.ApplyPreferences(pref)

	return ll
}
