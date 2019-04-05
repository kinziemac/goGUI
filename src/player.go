package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type player struct {
	id           int
	currentBlock int32
	score        int
	active       bool
}

func colorRect(renderer *sdl.Renderer, rect *sdl.Rect) {
	renderer.SetDrawColor(0, 0, 255, 255)
	renderer.FillRect(rect)
}

func newPlayer() (p player) {
	p.id = 1
	p.currentBlock = -1
	p.score = 0
	p.active = false

	return p
}

func (p *player) requestAccess() bool {

	return true
}
