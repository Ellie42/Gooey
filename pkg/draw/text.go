package draw

import (
	"git.agehadev.com/elliebelly/gooey/fonts"
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
	"git.agehadev.com/elliebelly/gooey/pkg/draw/font"
	"github.com/go-gl/gl/v4.6-compatibility/gl"
	"math"
)

var currentFont *font.Font
var textVAO uint32
var textVertVBO uint32
var textVertColourVBO uint32
var textUVVBO uint32
var fontTexture uint32

var textVertBuffer = make([]dimension.Vector3, 2048)
var textColourBuffer = make([]RGBA, 2048)
var textUVBuffer = make([]dimension.Vector2, 2048)

func Text(rect dimension.Rect, str string, sizePixels int) {
	SwitchProgram(programs[1])
	SwitchBlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	if currentFont == nil {
		currentFont = font.LoadFromHexString(fonts.SourceCodePro24)
	}

	gl.Enable(gl.TEXTURE_2D)

	if textVAO == 0 {
		textVAO = genVAO(1)[0]
		gl.BindVertexArray(textVAO)

		vbos := genVBO(3)
		textVertVBO, textVertColourVBO, textUVVBO = vbos[0], vbos[1], vbos[2]

		gl.EnableVertexAttribArray(0)
		gl.BindBuffer(gl.ARRAY_BUFFER, textVertVBO)
		gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

		gl.EnableVertexAttribArray(1)
		gl.BindBuffer(gl.ARRAY_BUFFER, textVertColourVBO)
		gl.VertexAttribPointer(1, 4, gl.FLOAT, false, 0, nil)

		gl.EnableVertexAttribArray(2)
		gl.BindBuffer(gl.ARRAY_BUFFER, textUVVBO)
		gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 0, nil)

		gl.GenTextures(1, &fontTexture)
		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, fontTexture)

		gl.TexImage2D(
			gl.TEXTURE_2D,
			0,
			gl.RGBA,
			int32(currentFont.BFF.ImageWidth),
			int32(currentFont.BFF.ImageHeight),
			0,
			gl.RGBA,
			gl.UNSIGNED_BYTE,
			gl.Ptr(currentFont.Data),
		)

		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_BASE_LEVEL, 0)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAX_LEVEL, 0)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
		//gl.TextureParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_LOD, 0)
		//gl.TextureParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAX_LOD, 0)
		//gl.TextureParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAX_LEVEL, 0)

		//gl.GenerateMipmap(gl.TEXTURE_2D)
		//gl.BindTexture(gl.TEXTURE_2D, 0)
	}

	resRatioX := float32(CurrentResolution.Width) / float32(CurrentResolution.Height)
	charHeight := currentFont.BFF.CellHeight
	cumStartPosX := rect.X + float32(0)

	desiredPixelHeight := float32(sizePixels)
	scaledCharHeight := 1 / float32(CurrentResolution.Height) * desiredPixelHeight

	perPixel := dimension.Vector2{
		1.0 / float32(CurrentResolution.Width),
		1.0 / float32(CurrentResolution.Height),
	}

	rect.Y = float32(math.Round(float64(rect.Y/perPixel.Y))) * perPixel.Y

	for i, c := range str {
		charWidth := currentFont.BFF.CharacterWidths[c]

		charHeightRelative := float32(1) * scaledCharHeight
		charWidthRelative := (float32(charWidth) / float32(charHeight)) * charHeightRelative / resRatioX

		cumStartPosX = float32(math.Round(float64(cumStartPosX/perPixel.X))) * perPixel.X

		copy(textVertBuffer[i*6:], preparePositionsForGL([]dimension.Vector3{
			{cumStartPosX, rect.Y + charHeightRelative, 1},
			{cumStartPosX, rect.Y, 1},
			{cumStartPosX + charWidthRelative, rect.Y, 1},

			{cumStartPosX, rect.Y + charHeightRelative, 1},
			{cumStartPosX + charWidthRelative, rect.Y, 1},
			{cumStartPosX + charWidthRelative, rect.Y + charHeightRelative, 1},
		}))

		uvs := currentFont.GetCharacterUV(uint8(c))

		copy(textUVBuffer[i*6:], []dimension.Vector2{
			{uvs.MinX, uvs.MinY},
			{uvs.MinX, uvs.MaxY},
			{uvs.MaxX, uvs.MaxY},

			{uvs.MinX, uvs.MinY},
			{uvs.MaxX, uvs.MaxY},
			{uvs.MaxX, uvs.MinY},
		})

		for j := i * 6; j < i*6+6; j++ {
			textColourBuffer[j] = White
		}

		cumStartPosX += charWidthRelative
	}

	gl.BindVertexArray(textVAO)

	gl.BindBuffer(gl.ARRAY_BUFFER, textVertVBO)
	gl.BufferData(gl.ARRAY_BUFFER, 4*3*len(textVertBuffer), gl.Ptr(textVertBuffer), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ARRAY_BUFFER, textVertColourVBO)
	gl.BufferData(gl.ARRAY_BUFFER, 4*4*len(textColourBuffer), gl.Ptr(textColourBuffer), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ARRAY_BUFFER, textUVVBO)
	gl.BufferData(gl.ARRAY_BUFFER, 4*2*len(textUVBuffer), gl.Ptr(textUVBuffer), gl.STATIC_DRAW)

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, fontTexture)

	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(textVertBuffer)/3))

	RestoreGLOptions()
}
