package widget

import "git.agehadev.com/elliebelly/gooey/pkg/dimension"

type PanelWidget struct {
	BaseWidget
	Content LinearLayout

	directionMultiplierX int
	directionMultiplierY int
}

func (p *PanelWidget) Render() {
	if p.Rect.X < 0 || p.Rect.X+p.Rect.Width > 1 {
		p.directionMultiplierX *= -1
	}

	if p.Rect.Y < 0 || p.Rect.Y+p.Rect.Height > 1 {
		p.directionMultiplierY *= -1
	}

	p.Rect.X += 0.001 * float32(p.directionMultiplierX)
	p.Rect.Y += 0.0015 * float32(p.directionMultiplierY)

	p.baseRender()
}

func NewPanel(widget ...Widget) *PanelWidget {
	pw := &PanelWidget{
		directionMultiplierX: 1,
		directionMultiplierY: 1,
	}

	pw.Rect = dimension.Rect{
		X:      0.25,
		Y:      0.25,
		Width:  0.5,
		Height: 0.5,
	}

	pw.Children = widget

	return pw
}
