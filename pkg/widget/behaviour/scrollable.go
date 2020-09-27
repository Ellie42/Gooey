package behaviour

import (
	"fmt"
	"git.agehadev.com/elliebelly/gooey/lib/eventmanager"
)

type Scrollable struct {
	RectalBehaviour
}

func (s *Scrollable) Init(eventManager *eventmanager.EventManager) {
	s.eventManager = eventManager
	s.inputHandle = eventManager.RegisterMouseScrollHandler(s)
}

func (s Scrollable) GetZIndex() int {
	return 0
}

func (s Scrollable) OnScroll(xoff, yoff float64) {
	fmt.Printf("Scroll me baby %f %f\n", xoff, yoff)
}
