package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	renderWidth  = 1920
	renderHeight = 1080
	windowWidth  = 1280
	windowHeight = 720
)

type Game struct {
	ticks int
}

func (g *Game) Update() error {
	g.ticks++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 22, G: 26, B: 34, A: 255})
	ebitenutil.DebugPrint(screen, "ebi-mawaru\nRender: 1920x1080\nWindow: 1280x720 resizable")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return renderWidth, renderHeight
}

func main() {
	ebiten.SetWindowTitle("ebi-mawaru")
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
