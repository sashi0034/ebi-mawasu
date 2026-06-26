package main

import (
	"ebi-mawaru/cmd/game/asset"
	"ebi-mawaru/cmd/game/core"
	"ebi-mawaru/cmd/game/gameplay"
	"ebi-mawaru/cmd/game/utils"
)

const (
	renderWidth  = 1920
	renderHeight = 1080
)

type Game struct {
	actors      core.ActorList
	windowState *utils.WindowStateManager
}

func newGame(windowState *utils.WindowStateManager) (*Game, error) {
	assets, err := asset.NewAsset()
	if err != nil {
		return nil, err
	}

	game := &Game{
		windowState: windowState,
	}
	game.actors.Add(gameplay.NewGameplayScene(assets))

	return game, nil
}
