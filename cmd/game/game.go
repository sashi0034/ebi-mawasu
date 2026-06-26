package main

import (
	"ebi-mawaru/cmd/game/asset"
	"ebi-mawaru/cmd/game/utils"
)

const (
	renderWidth  = 1920
	renderHeight = 1080
)

type Game struct {
	ticks       int
	assets      *asset.Asset
	windowState *utils.WindowStateManager
}

func newGame(windowState *utils.WindowStateManager) (*Game, error) {
	assets, err := asset.NewAsset()
	if err != nil {
		return nil, err
	}

	return &Game{
		assets:      assets,
		windowState: windowState,
	}, nil
}
