package dimension

type Vector2 struct {
	X, Y float32
}

func (v Vector2) Sub(position Vector2) Vector2 {
	v.X -= position.X
	v.Y -= position.Y
	return v
}
