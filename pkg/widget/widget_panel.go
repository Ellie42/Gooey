package widget

import (
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

	p.RenderStyles(
		p.GetRectAbsolute(),
		p.Prefs.StyleSettings,
	)
	p.RenderChildren()
	p.ShowBaseDebug()
}

func NewPanel(pref *settings.WidgetPreferences, widget ...Widget) *PanelWidget {
	pw := &PanelWidget{}

	pw.draggable.DragRect(pw.GetRectAbsolute)

	pw.AddChildWithParent(pw, widget...)
	pw.ApplyPreferences(pref)

	return pw
}
