package behaviour

import (
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
	"git.agehadev.com/elliebelly/gooey/lib/eventmanager"
)

type BehaviourSet struct {
	*Draggable
	*Clickable
	*Scrollable
}

func (b *BehaviourSet) Init(manager *eventmanager.EventManager, rectProviders ...func() dimension.Rect) {
	if b.Draggable != nil {
		b.Draggable.Init(manager)
		b.Draggable.AddRectProviders(rectProviders...)
	}

	if b.Clickable != nil {
		b.Clickable.Init(manager)
		b.Clickable.AddRectProviders(rectProviders...)
	}

	if b.Scrollable != nil {
		b.Scrollable.Init(manager)
		b.Scrollable.AddRectProviders(rectProviders...)
	}
}
