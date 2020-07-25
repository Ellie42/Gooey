package dimension

type DimensionsInt struct {
	Width, Height int
}

type DimensionsFloat32 struct {
	Width, Height float32
}

type Dimensions struct {
	Width, Height *Size
}
