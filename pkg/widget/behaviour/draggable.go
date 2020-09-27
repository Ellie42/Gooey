package behaviour

import (
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
	"git.agehadev.com/elliebelly/gooey/lib/eventmanager"
)

type Draggable struct {
	RectalBehaviour

	isDragging         bool
	lastCursorPosition dimension.Vector2
}

func (d *Draggable) Init(eventManager *eventmanager.EventManager) {
	d.eventManager = eventManager
	d.inputHandle = eventManager.RegisterMouseClickHandler(d)
}

func (d *Draggable) GetZIndex() int {
	return 0
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

func (d *Draggable) GetDragDiff() dimension.Vector2 {
	currentPosition := d.eventManager.MousePosition()

	diff := currentPosition.Sub(d.lastCursorPosition)

	d.lastCursorPosition = currentPosition

	return d.eventManager.PixelPositionToRelative(diff)
}
