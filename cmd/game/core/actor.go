package core

import "github.com/hajimehoshi/ebiten/v2"

type Actor interface {
	Update() error

	Draw(screen *ebiten.Image)

	Alive() bool
}

type UpdatePrioritized interface {
	UpdatePriority() int
}

type DrawPrioritized interface {
	DrawPriority() int
}
