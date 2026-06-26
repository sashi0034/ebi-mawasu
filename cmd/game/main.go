package main

import (
	"bytes"
	"fmt"
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	renderWidth  = 1920
	renderHeight = 1080
	windowWidth  = 1280
	windowHeight = 720
)

type Game struct {
	ticks      int
	textFace   text.Face
	iconImage  *ebiten.Image
	fontSource string
}

func newGame() (*Game, error) {
	notoFont, fontSource, err := loadFont([]string{
		"assets/fonts/NotoSans/NotoSansMonoCJKjp-Bold.otf",
		"assets/fonts/NotoSans/NotoSansMonoCJKtc-Bold.otf",
		// "assets/fonts/NotoSans/NotoSansMonoCJKsc-Bold.otf",
		// "assets/fonts/NotoSans/NotoSansMonoCJKkr-Bold.otf",
	}, "Noto Sans Mono CJK")
	if err != nil {
		return nil, err
	}

	iconFont, _, err := loadFont([]string{
		"assets/fonts/materialdesignicons/materialdesignicons-webfont.ttf",
	}, "Material Design Icons")
	if err != nil {
		return nil, err
	}

	notoFace := &text.GoTextFace{
		Source: notoFont,
		Size:   56,
	}
	iconFace := &text.GoTextFace{
		Source: iconFont,
		Size:   64,
	}
	textFace, err := text.NewMultiFace(notoFace, iconFace)
	if err != nil {
		return nil, fmt.Errorf("create text fallback face: %w", err)
	}

	return &Game{
		textFace:   textFace,
		iconImage:  renderIconImage('\U000F02DC', iconFace, 160),
		fontSource: fontSource,
	}, nil
}

func loadFont(paths []string, name string) (*text.GoTextFaceSource, string, error) {
	for _, path := range paths {
		b, actualPath, err := readAsset(path)
		if err != nil {
			continue
		}

		source, err := text.NewGoTextFaceSource(bytes.NewReader(b))
		if err != nil {
			return nil, "", fmt.Errorf("parse %s at %s: %w", name, actualPath, err)
		}
		return source, actualPath, nil
	}
	return nil, "", fmt.Errorf("load %s: none of %v found", name, paths)
}

func readAsset(path string) ([]byte, string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, "", fmt.Errorf("read asset %s from working directory: %w", path, err)
	}
	return b, path, nil
}

func (g *Game) Update() error {
	g.ticks++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 22, G: 26, B: 34, A: 255})
	drawText(screen, "これは日本語です "+string('\U000F02DC'), g.textFace, 96, 140, color.RGBA{R: 240, G: 244, B: 255, A: 255})
	drawText(screen, "NotoSansMonoCJKjp-Bold + Material Design Icons fallback", g.textFace, 96, 230, color.RGBA{R: 168, G: 176, B: 202, A: 255})
	drawText(screen, g.fontSource, g.textFace, 96, 300, color.RGBA{R: 126, G: 138, B: 170, A: 255})
	g.drawIconImage(screen, 160, 500, 1.4, color.RGBA{R: 110, G: 231, B: 183, A: 255})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return renderWidth, renderHeight
}

func main() {
	game, err := newGame()
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowTitle("ebi-mawaru")
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func renderIconImage(r rune, face text.Face, size int) *ebiten.Image {
	img := ebiten.NewImage(size, size)
	drawText(img, string(r), face, 32, 32, color.White)
	return img
}

func (g *Game) drawIconImage(screen *ebiten.Image, x, y, scale float64, tint color.RGBA) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleWithColor(tint)
	screen.DrawImage(g.iconImage, op)
}

func drawText(dst *ebiten.Image, s string, face text.Face, x, y int, c color.Color) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	op.ColorScale.ScaleWithColor(c)
	text.Draw(dst, s, face, op)
}
