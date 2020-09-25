package font

import (
	"git.agehadev.com/elliebelly/gooey/lib/binary"
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
)

type BFF struct {
	ImageWidth          uint32
	ImageHeight         uint32
	CellWidth           uint32
	CellHeight          uint32
	BPP                 uint8
	ASCIIStartCharacter uint8
	CharacterWidths     []uint8
}

func (b BFF) GetCharWidth(ascii uint8) uint8 {
	return b.CharacterWidths[ascii]
}

type Font struct {
	Data []byte
	BFF  *BFF

	charactersPerRow int
	numRows          int
}

func (b *Font) Parse(str string) error {
	data, err := decompressFontHex(str)

	if err != nil {
		return err
	}

	b.BFF = parseBFFHeader(data[:0x115])

	b.Data = data[0x115:]

	b.charactersPerRow = int(b.BFF.ImageWidth) / int(b.BFF.CellWidth)
	b.numRows = int(b.BFF.ImageHeight) / int(b.BFF.CellHeight)

	return nil
}

func parseBFFHeader(bytes []byte) *BFF {
	br := binary.ByteReader{
		Data: bytes,
	}

	br.ReadUint8()
	br.ReadUint8()

	return &BFF{
		br.ReadUint32(),
		br.ReadUint32(),
		br.ReadUint32(),
		br.ReadUint32(),
		br.ReadUint8(),
		br.ReadUint8(),
		br.Remaining(),
	}
}

func (b Font) GetCharacterUV(ascii uint8) dimension.BoundingBox {
	offset := int(ascii) - int(b.BFF.ASCIIStartCharacter)

	row := offset / b.charactersPerRow
	col := offset % b.charactersPerRow

	colFactor := float32(b.BFF.CellWidth) / float32(b.BFF.ImageWidth)
	rowFactor := float32(b.BFF.CellHeight) / float32(b.BFF.ImageHeight)

	width := float32(b.BFF.GetCharWidth(ascii)) / float32(b.BFF.CellWidth)

	return dimension.BoundingBox{
		float32(col) * colFactor,
		float32(row) * rowFactor,
		float32(col)*colFactor + colFactor - ((1 - width) * colFactor),
		float32(row)*rowFactor + rowFactor,
	}
}
