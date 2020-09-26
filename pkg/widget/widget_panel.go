package widget

import (
	"git.agehadev.com/elliebelly/gooey/pkg/widget/settings"
)

type PanelWidget struct {
	BaseWidget
	Content   LinearLayout
}

func (p *PanelWidget) Render() {
	if p.Behaviours.Draggable != nil && p.Behaviours.Draggable.IsDragging() {
		absRect := p.GetRectAbsolute()
		dragDist := p.Behaviours.Draggable.GetDragDiff()
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

	pw.AddChildWithParent(pw, widget...)
	pw.ApplyPreferences(pref)

	return pw
}
