package main

import (
	"github.com/wmattei/open-cloud-game/pkg/engine"
)

type FpsGame struct{}

func (f *FpsGame) Update() {}
func (f *FpsGame) Render() {}

func main() {
	engine := engine.NewEngine()
	engine.Run(&FpsGame{}, 800, 600)

}
