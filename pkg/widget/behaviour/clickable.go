package behaviour

import (
	"fmt"
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
	"git.agehadev.com/elliebelly/gooey/lib/eventmanager"
)

type Clickable struct {
	RectalBehaviour
}

func (c *Clickable) Init(manager *eventmanager.EventManager) {
	c.eventManager = manager
	c.inputHandle = manager.RegisterMouseClickHandler(c)
}

func (c Clickable) GetZIndex() int {
	return 0
}

func (c Clickable) OnClick(position dimension.Vector2) {
	fmt.Println("OH GOD IT HURTS")
}

func (c Clickable) OnMouseUp() {
}
