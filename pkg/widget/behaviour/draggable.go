package behaviour

import (
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
	"git.agehadev.com/elliebelly/gooey/lib/eventmanager"
)

type Draggable struct {
	DragRects []func() dimension.Rect

	isDragging         bool
	inputHandle        int
	lastCursorPosition dimension.Vector2
	eventManager       *eventmanager.EventManager
}

func (d *Draggable) Init(eventManager *eventmanager.EventManager) {
	d.eventManager = eventManager
	d.inputHandle = eventManager.RegisterRect(d)
}

func (d *Draggable) GetZIndex() int {
	return 0
}

func (d *Draggable) GetRects() []dimension.Rect {
	rects := make([]dimension.Rect, len(d.DragRects))

	for _, rectFunc := range d.DragRects {
		rects = append(rects, rectFunc())
	}

	return rects
}

func (d *Draggable) OnMouseUp() {
	d.isDragging = false
}

func (d *Draggable) OnClick(position dimension.Vector2) {
	d.isDragging = true
	d.lastCursorPosition = position
}

func (d Draggable) IsDragging() bool {
	return d.isDragging
}

func (d *Draggable) DragRect(rect ...func() dimension.Rect) {
	d.DragRects = rect
}

func (d *Draggable) GetDragDiff() dimension.Vector2 {
	currentPosition := d.eventManager.MousePosition()

	diff := currentPosition.Sub(d.lastCursorPosition)

	d.lastCursorPosition = currentPosition

	return d.eventManager.PixelPositionToRelative(diff)
}
