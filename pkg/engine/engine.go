package engine

import (
	"fmt"
	"math"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	errr "github.com/wmattei/open-cloud-game/pkg/error"
	"github.com/wmattei/open-cloud-game/pkg/mymath"
)

type Engine struct{}

func NewEngine() *Engine {
	return &Engine{}
}

func init() {
	runtime.LockOSThread()
	err := glfw.Init()
	errr.PanicIfError(err)

	glfw.WindowHint(glfw.DoubleBuffer, glfw.True)
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
}

func (e *Engine) Run(g Game, width, height int) error {
	window, err := glfw.CreateWindow(width, height, "HEEEEEE", nil, nil)
	if err != nil {
		return err
	}
	window.MakeContextCurrent()

	program := initGl()
	gl.UseProgram(program)
	gl.Viewport(0, 0, int32(width), int32(height))

	camera := NewPerspectiveCamera(
		[3]float32{0, 1, 0},
		[3]float32{0, 1, 0},
		0,
		0,
		math.Pi/2,
		float32(width)/float32(height),
		0.01,
		1000,
	)

	scene := NewScene()
	scene.entities = append(scene.entities, NewEntity())

	for !window.ShouldClose() {
		g.Update()
		g.Render()

		view := camera.GetViewMatrix()
		viewFlatten := view.Flatten()
		projection := camera.GetProjectionMatrix()
		projectionFlatten := projection.Flatten()

		model := mymath.GetIdentityMatrix()
		modelFlatten := model.Flatten()

		fmt.Println(modelFlatten)

		viewLoc := gl.GetUniformLocation(program, gl.Str("view\x00"))
		projLoc := gl.GetUniformLocation(program, gl.Str("projection\x00"))
		modelLoc := gl.GetUniformLocation(program, gl.Str("model\x00"))

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT | gl.STENCIL_BUFFER_BIT)

		gl.UniformMatrix4fv(viewLoc, 1, false, &viewFlatten[0])
		gl.UniformMatrix4fv(projLoc, 1, false, &projectionFlatten[0])
		gl.UniformMatrix4fv(modelLoc, 1, false, &modelFlatten[0])

		scene.Render()

		window.SwapBuffers()
		glfw.PollEvents()
	}

	return nil
}
