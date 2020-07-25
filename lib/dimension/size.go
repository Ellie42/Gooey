package dimension

type SizeUnit int

const (
	SizeUnitRatio SizeUnit = iota
	SizeUnitPixels
)

type Size struct {
	Amount float32
	Unit   SizeUnit
}