package dimension

type SizeUnit int

const (
	SizeUnitPixels SizeUnit = iota
	SizeUnitRatio
)

type Size struct {
	Amount float32
	Unit   SizeUnit
}
