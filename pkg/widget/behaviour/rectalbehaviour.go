package behaviour

import "git.agehadev.com/elliebelly/gooey/lib/dimension"

type RectalBehaviour struct {
	Rects   []func() dimension.Rect
}

func (d *RectalBehaviour) GetRects() []dimension.Rect {
	rects := make([]dimension.Rect, 0)

	for _, rectFunc := range d.Rects {
		rects = append(rects, rectFunc())
	}

	return rects
}

func (d *RectalBehaviour) AddRectProviders(rect ...func() dimension.Rect) {
	d.Rects = rect
}
