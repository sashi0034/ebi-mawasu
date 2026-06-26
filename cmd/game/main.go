package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"ebi-mawaru/cmd/game/mdi"
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

	g.ticks++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	textFace := g.assets.TextFace()

	screen.Fill(color.RGBA{R: 22, G: 26, B: 34, A: 255})
	drawText(screen, "これは日本語です "+string(mdi.Home), textFace, 96, 140, color.RGBA{R: 240, G: 244, B: 255, A: 255})
	drawText(screen, "NotoSansMonoCJKjp-Bold + Material Design Icons fallback", textFace, 96, 230, color.RGBA{R: 168, G: 176, B: 202, A: 255})
	drawText(screen, g.assets.FontSource(), textFace, 96, 300, color.RGBA{R: 126, G: 138, B: 170, A: 255})

	g.drawIconImage(screen, 160, 500, 1.4, color.RGBA{R: 110, G: 231, B: 183, A: 255})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return renderWidth, renderHeight
}

func (g *Game) drawIconImage(screen *ebiten.Image, x, y, scale float64, tint color.RGBA) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleWithColor(tint)
	screen.DrawImage(g.assets.IconImage(mdi.Home), op)
}

func drawText(dst *ebiten.Image, s string, face text.Face, x, y int, c color.Color) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	op.ColorScale.ScaleWithColor(c)
	text.Draw(dst, s, face, op)
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
