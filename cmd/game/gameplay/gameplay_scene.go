package gameplay

import (
	"github.com/hajimehoshi/ebiten/v2"

	"ebi-mawaru/cmd/game/asset"
	"ebi-mawaru/cmd/game/core"
)

type GameplayScene struct {
	children core.ActorList
}

var _ core.Actor = (*GameplayScene)(nil) // implement Actor

func NewGameplayScene(assets *asset.Asset) *GameplayScene {
	scene := &GameplayScene{}
	scene.children.Add(NewMyObject(assets))
	return scene
}

func (s *GameplayScene) Update() error {
	return s.children.Update()
}

func (s *GameplayScene) Draw(screen *ebiten.Image) {
	s.children.Draw(screen)
}

func (s *GameplayScene) Alive() bool {
	return true
}
