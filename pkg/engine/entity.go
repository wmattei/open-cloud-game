package engine

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Entity struct {
	vao uint32
	vbo uint32
	ebo uint32

	shouldRender bool
}

func NewEntity() *Entity {
	entity := &Entity{}

	gl.GenVertexArrays(1, &entity.vao)
	gl.BindVertexArray(entity.vao)

	vertices := entity.GetVertices()
	indices := entity.GetIndices()

	gl.GenBuffers(1, &entity.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, entity.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.GenBuffers(1, &entity.ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, entity.ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 7*4, gl.PtrOffset(0))

	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointer(1, 4, gl.FLOAT, false, 7*4, gl.PtrOffset(3*4))

	gl.BindVertexArray(0)

	entity.shouldRender = true

	return entity
}

func (e *Entity) Render() {
	if e.shouldRender {
		indexCount := len(e.GetIndices())
		gl.BindVertexArray(e.vao)
		gl.DrawElements(gl.TRIANGLES, int32(indexCount), gl.UNSIGNED_INT, gl.PtrOffset(0))
		gl.BindVertexArray(0)
	}
}

func (e *Entity) GetVertices() []float32 {
	return []float32{
		-0.5, -0.5, -0.5, 1.0, 0.0, 0.0, 1.0, // Vertex 0: Red
		0.5, -0.5, -0.5, 0.0, 1.0, 0.0, 1.0, // Vertex 1: Green
		0.5, 0.5, -0.5, 0.0, 0.0, 1.0, 1.0, // Vertex 2: Blue
		-0.5, 0.5, -0.5, 1.0, 1.0, 0.0, 1.0, // Vertex 3: Yellow
		-0.5, -0.5, 0.5, 1.0, 0.0, 1.0, 1.0, // Vertex 4: Magenta
		0.5, -0.5, 0.5, 0.0, 1.0, 1.0, 1.0, // Vertex 5: Cyan
		0.5, 0.5, 0.5, 1.0, 0.5, 0.0, 1.0, // Vertex 6: Orange
		-0.5, 0.5, 0.5, 0.5, 0.0, 1.0, 1.0, // Vertex 7: Purple
	}
}

func (e *Entity) GetIndices() []uint32 {
	return []uint32{
		0, 1, 2,
		2, 3, 0,
		4, 5, 6,
		6, 7, 4,
		0, 4, 7,
		7, 3, 0,
		1, 5, 6,
		6, 2, 1,
		3, 7, 6,
		6, 2, 3,
		0, 4, 5,
		5, 1, 0,
	}
}
