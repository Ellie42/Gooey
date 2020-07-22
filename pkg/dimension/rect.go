package dimension

type Rect struct {
	X, Y, Width, Height float32
}

func (r Rect) RelativeTo(parent Rect) *Rect {
	r.X = parent.X + r.X*float32(parent.Width)
	r.Y = parent.Y + r.Y*float32(parent.Height)
	r.Width = r.Width * float32(parent.Width)
	r.Height = r.Height * float32(parent.Height)
	return &r
}
