package dimension

type Vector2 struct {
	X, Y float32
}

func (v Vector2) Sub(position Vector2) Vector2 {
	v.X -= position.X
	v.Y -= position.Y
	return v
}

type Vector3 struct {
	X, Y, Z float32
}

func (v Vector3) Sub(position Vector3) Vector3 {
	v.X -= position.X
	v.Y -= position.Y
	v.Z -= position.Z
	return v
}
