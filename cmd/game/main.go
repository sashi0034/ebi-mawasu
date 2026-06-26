package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"ebi-mawaru/cmd/game/utils"
)

const (
	windowWidth  = 1280
	windowHeight = 720
)

func (g *Game) Update() error {
	if ebiten.IsWindowBeingClosed() {
		_ = g.windowState.Save()
		return ebiten.Termination
	}

	return g.actors.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.actors.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return renderWidth, renderHeight
}

func main() {
	windowState := utils.NewWindowStateManager("ebi-mawaru")
	game, err := newGame(windowState)
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowTitle("ebi-mawaru")
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowClosingHandled(true)

	// Air によるリロード時にウィンドウ位置を復元したい
	windowState.RestoreIfNeeded()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
