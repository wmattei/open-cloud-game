package engine

type Scene struct {
	entities []*Entity
}

func NewScene() *Scene {
	return &Scene{}
}

func (s *Scene) Render() {
	for _, e := range s.entities {
		e.Render()
	}
}
