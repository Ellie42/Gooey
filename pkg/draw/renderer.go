package draw

import (
	"fmt"
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
	"git.agehadev.com/elliebelly/gooey/lib/shader"
	"github.com/go-gl/gl/v4.6-compatibility/gl"
	"github.com/go-gl/mathgl/mgl32"
)

var CurrentResolution dimension.DimensionsInt

const (
	colouredVertShader = `
    #version 460

	uniform mat4 projection;

    in vec3 vpos;
	in vec4 vcolour;
	in vec2 uv;

	out vec4 _vcolour;
	out vec2 _uv;

    void main() {
        gl_Position = vec4(vpos, 1.0);
		_vcolour = vcolour;
		_uv = uv;
    }
` + "\x00"

	colouredFragShader = `
	#version 460

	in vec4 _vcolour;

	out vec4 outputColor;

	void main() {
		outputColor = _vcolour;
	}
` + "\x00"

	texturedFragShader = `
	#version 460

	in vec4 _vcolour;
	in vec2 _uv;

	uniform sampler2D tex;

	out vec4 outputColor;

	void main() {
		outputColor = _vcolour * vec4(1.0, 1.0, 1.0, texture2D(tex, _uv).r);
	}
` + "\x00"
)

var currentProgram uint32
var blendOptions map[uint32]uint32
var programs []uint32

type Renderer struct {
}

func (r *Renderer) Init() {
	if err := gl.Init(); err != nil {
		panic(fmt.Sprintf("failed to initialise opengl"))
	}

	programs = make([]uint32, 0)

	programs = append(programs, shader.CompileProgram(colouredVertShader, colouredFragShader))
	programs = append(programs, shader.CompileProgram(colouredVertShader, texturedFragShader))

	gl.ClearColor(0, 0, 0, 1.0)

	blendOptions = make(map[uint32]uint32)

	blendOptions[gl.SRC_ALPHA] = gl.ONE_MINUS_SRC_ALPHA

	gl.UseProgram(programs[0])

	loc := gl.GetUniformLocation(programs[1], gl.Str("projection\x00"))
	orthMatrix := mgl32.Ortho(0, 0.9, 0, 1, 0, 100)

	if loc != -1 {
		gl.UniformMatrix4fv(loc, 1, false, &orthMatrix[0])
	}

	RestoreGLOptions()
}

func (r *Renderer) Clear(res dimension.DimensionsInt) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.UseProgram(programs[0])
}

func SwitchProgram(program uint32) {
	currentProgram = program
	gl.UseProgram(program)
}

func RestoreGLOptions() {
	currentProgram = programs[0]
	gl.UseProgram(programs[0])

	for option, value := range blendOptions {
		gl.BlendFunc(option, value)
	}
}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func SwitchBlendFunc(funcID uint32, option uint32) {
	gl.Enable(gl.BLEND)
	gl.BlendFunc(funcID, option)
}
