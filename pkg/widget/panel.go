package widget

import (
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
	"git.agehadev.com/elliebelly/gooey/pkg/widget/behaviour"
	"git.agehadev.com/elliebelly/gooey/pkg/widget/settings"
)

type PanelWidget struct {
	BaseWidget
	Content   LinearLayout
	draggable behaviour.Draggable
}

func (p *PanelWidget) Init() {
	p.draggable.Init(&Context.EventManager)
	p.InitChildren(p)
}

func (p *PanelWidget) Render() {
	if p.draggable.IsDragging() {
		absRect := p.GetRectAbsolute()
		dragDist := p.draggable.GetDragDiff()
		p.Rect.X += dragDist.X * (1 / absRect.Width)
		p.Rect.Y += dragDist.Y * (1 / absRect.Height)
	}

	p.RenderChildren()
	p.ShowBaseDebug()
}

func NewPanel(pref *settings.WidgetPreferences, widget ...Widget) *PanelWidget {
	pw := &PanelWidget{}

	pw.draggable.DragRect(pw.GetRectAbsolute)

	pw.Rect = &dimension.Rect{
		X:      0,
		Y:      0,
		Width:  1,
		Height: 1,
	}

	pw.Children = widget

	pw.ApplyPreferences(pref)

	return pw
}
