package draw

import (
	"encoding/hex"
	"errors"
)

type RGBA struct {
	R, G, B, A float32
}

func (r *RGBA) SetHexRGBA(val string) (err error) {
	if len(val) < 6 {
		return errors.New("string must include RGB[A] hex segments")
	}

	r.R = mustGetFloatFromHex(string([]byte(val)[0:2]))
	r.G = mustGetFloatFromHex(string([]byte(val)[2:4]))
	r.B = mustGetFloatFromHex(string([]byte(val)[4:6]))

	if len(val) == 8 {
		r.A = mustGetFloatFromHex(string([]byte(val)[6:8]))
	} else {
		r.A = 1
	}

	return nil
}

func mustGetFloatFromHex(str string) float32 {
	var byteVal []byte

	byteVal, err := hex.DecodeString(str)

	if err != nil {
		panic(err)
	}

	return float32(byteVal[0]) / float32(0xFF)
}

var (
	Red   RGBA = RGBA{1, 0, 0, 1}
	Green RGBA = RGBA{0, 1, 0, 1}
	Blue  RGBA = RGBA{0, 0, 1, 1}
	White RGBA = RGBA{1, 1, 1, 1}
)

func NewRGBAFromHex(h string) RGBA {
	rgba := RGBA{}

	if err := rgba.SetHexRGBA(h); err != nil {
		panic(err)
	}

	return rgba
}
