package engine

import "github.com/go-gl/gl/v4.1-core/gl"

func initGl() uint32 {
	gl.Init()
	program := gl.CreateProgram()
	initShaders(program)
	gl.LinkProgram(program)

	gl.Enable(gl.DEPTH_TEST)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := make([]byte, logLength)
		gl.GetProgramInfoLog(program, logLength, nil, &log[0])
		panic("failed to link program:" + string(log))
	}

	destroyShaders()

	return program
}
