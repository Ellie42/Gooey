package widget

type Widget interface {
	SetIndex(index int)
	Init(parent WidgetParent)
	Render()
}
