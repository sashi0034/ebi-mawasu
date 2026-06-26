package main

import (
	"ebi-mawaru/cmd/game/asset"
)

const (
	renderWidth  = 1920
	renderHeight = 1080
)

type Game struct {
	ticks  int
	assets *asset.Asset
}

func newGame() (*Game, error) {
	assets, err := asset.NewAsset()
	if err != nil {
		return nil, err
	}

	return &Game{
		assets: assets,
	}, nil
}
