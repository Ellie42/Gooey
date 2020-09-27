package widget

import (
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
	"git.agehadev.com/elliebelly/gooey/pkg/widget/settings"
)

type LinearLayoutAlignment int

const (
	LLAlignmentHorizontal LinearLayoutAlignment = iota
	LLAlignmentVertical
)

type LinearLayout struct {
	BaseWidget
	Alignment     LinearLayoutAlignment
	Initialised   bool
	FitToChildren bool
	stepOffset    float32
}

func (l *LinearLayout) RecalculateChildRects() {
	if l.Children == nil {
		return
	}

	parentRect := l.GetRectAbsolute()

	if l.Prefs.Padding != nil {
		parentRect = parentRect.WithPaddingAbsolute(l.Prefs.Padding.ToDirectionalRect(Context.Resolution))
	}

	step := float32(1) / float32(len(l.Children))
	stepOffset := float32(0)

	stepGroups := make(map[int]float32)
	groupSteps := make(map[int]float32)

	gIndex := 0
	childrenInGroup := 0

	for i, c := range l.Children {
		var x, y, w, h float32 = 0, 0, 1, 1

		if l.Alignment == LLAlignmentHorizontal {
			x = stepOffset
			w = step
		} else {
			y = stepOffset
			h = step
		}

		childRect := dimension.Rect{
			X:      x,
			Y:      y,
			Width:  w,
			Height: h,
		}

		cPrefs := c.GetPrefs()

		childFinalStepSize := step

		if cPrefs.DimensionBounds != nil {
			cPrefs.DimensionBounds.Bind(&childRect, parentRect, dimension.Width, Context.Resolution)
			cPrefs.DimensionBounds.Bind(&childRect, parentRect, dimension.Height, Context.Resolution)

			if l.Alignment == LLAlignmentHorizontal {
				childFinalStepSize = childRect.Width
			} else {
				childFinalStepSize = childRect.Height
			}

			diff := step - childFinalStepSize

			groupSteps[gIndex] += diff

			stepGroups[i] = childFinalStepSize

			if childrenInGroup > 0 {
				groupSteps[gIndex] /= float32(childrenInGroup)
				gIndex++
				childrenInGroup = 0
			}
		} else {
			groupSteps[gIndex] += childFinalStepSize
			childrenInGroup++
		}
	}

	if childrenInGroup > 0 {
		groupSteps[gIndex] /= float32(childrenInGroup)
	}

	stepOffset = 0
	gIndex = 0

	steps := make([][]float32, len(l.Children))

	for i, _ := range l.Children {
		stepSize := float32(1)

		if childFixedSize, ok := stepGroups[i]; ok {
			stepSize = childFixedSize
		} else {
			stepSize = groupSteps[gIndex]

			if _, ok := stepGroups[i+1]; ok {
				gIndex++
			}
		}

		steps[i] = []float32{stepOffset, stepSize}
		stepOffset += stepSize
	}

	for i := 0; i < len(l.Children); i++ {
		var x, y, w, h float32 = 0, 0, 1, 1

		if l.Alignment == LLAlignmentHorizontal {
			x = steps[i][0]
			w = steps[i][1]
		} else {
			y = steps[i][0]
			h = steps[i][1]
		}

		childRect := dimension.Rect{
			X: x,
			Y: y,
			//TODO fix this when using > 2 children, it will surely overshoot
			Width:  w,
			Height: h,
		}.RelativeToAbsolute(l.GetRectAbsolute())

		siblingRect := childRect.RelativeWithin(parentRect)

		if l.Alignment == LLAlignmentHorizontal {
			stepOffset += siblingRect.Width
		} else {
			stepOffset += siblingRect.Height
		}

		l.childRects[i] = childRect
	}

	for _, c := range l.Children {
		c.RecalculateChildRects()
	}
}

//func (l *LinearLayout) GetChildRectAbsolute(index int) dimension.Rect {
//	parentRect := l.GetRectAbsolute()
//
//	if l.Prefs.Padding != nil {
//		parentRect = parentRect.WithPaddingAbsolute(l.Prefs.Padding.ToDirectionalRect(Context.Resolution))
//	}
//
//	stepOffset := float32(0)
//
//	for i := 0; i < index; i++ {
//		child := l.Children[i]
//		siblingRect := child.GetRectAbsolute().RelativeWithin(parentRect)
//		if l.Alignment == LLAlignmentHorizontal {
//			stepOffset += siblingRect.Width
//		} else {
//			stepOffset += siblingRect.Height
//		}
//	}
//
//	remaining := 1 - stepOffset
//	remainingForChild := remaining / float32(len(l.Children)-index)
//
//	return l.getChildRectAbsolute(index, remainingForChild, stepOffset)
//}
//
//func (l *LinearLayout) getChildRectAbsolute(index int, step float32, stepOffset float32) dimension.Rect {
//	var x, y, w, h float32 = 0, 0, 1, 1
//
//	if l.Alignment == LLAlignmentHorizontal {
//		x = stepOffset
//		w = step
//	} else {
//		y = stepOffset
//		h = step
//	}
//
//	childRect := dimension.Rect{
//		X: x,
//		Y: y,
//		//TODO fix this when using > 2 children, it will surely overshoot
//		Width:  w,
//		Height: h,
//	}.RelativeToAbsolute(l.GetRectAbsolute())
//
//	if l.Prefs.Padding != nil {
//		childPadding := *l.Prefs.Padding
//
//		if index > 0 {
//			// Remove padding for all after first so only outer edge has padding
//			if l.Alignment == LLAlignmentHorizontal {
//				childPadding.Left.Amount = 0
//			} else {
//				childPadding.Bottom.Amount = 0
//			}
//		}
//
//		childRect = childRect.WithPaddingRelative(childPadding.ToDirectionalRect(Context.Resolution))
//	}
//	return childRect
//}

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
