package dimension

type Rect struct {
	X, Y, Width, Height float32
}

func (r Rect) RelativeTo(parent Rect) *Rect {
	return &Rect{
		X:      r.X * float32(parent.Width),
		Y:      r.Y * float32(parent.Height),
		Width:  r.Width * float32(parent.Width),
		Height: r.Height * float32(parent.Height),
	}
}
