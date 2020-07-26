package renderer

import (
	"fmt"
	"git.agehadev.com/elliebelly/gooey/lib/shader"
	"github.com/go-gl/gl/v4.6-core/gl"
)

const (
	vertexShaderSource = `
    #version 460

    in vec3 vpos;
	//in vec4 vcolour;
	layout(location = 1) in vec4 vcolour;

	out vec4 _vcolour;

    void main() {
        gl_Position = vec4(vpos, 1.0);
		_vcolour = vcolour;
    }
` + "\x00"

	fragmentShaderSource = `
	#version 460

	in vec4 _vcolour;

	out vec4 outputColor;

	void main() {
		//outputColor = vec4(0.1,1,0.2,1);
		outputColor = _vcolour;
	}
` + "\x00"
)

type Renderer struct {
	programs       []uint32
	currentProgram uint32
}

func (r *Renderer) Init() {
	if err := gl.Init(); err != nil {
		panic(fmt.Sprintf("failed to initialise opengl"))
	}

	r.programs = make([]uint32, 0)

	r.programs = append(r.programs, shader.CompileProgram(vertexShaderSource, fragmentShaderSource))

	gl.ClearColor(0, 0, 0, 1.0)
}

func (r *Renderer) Clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.UseProgram(r.programs[0])
}

func (r *Renderer) SwitchProgram(program uint32) {
	r.currentProgram = program
	gl.UseProgram(program)
}

func (r *Renderer) RestoreProgram() {
	r.currentProgram = r.programs[0]
	gl.UseProgram(r.programs[0])
}

func NewRenderer() *Renderer {
	return &Renderer{}
}
