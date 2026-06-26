package gameplay

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"ebi-mawaru/cmd/game/asset"
	"ebi-mawaru/cmd/game/core"
	"ebi-mawaru/cmd/game/mdi"
)

type MyObject struct {
	assets *asset.Asset
	ticks  int
}

var _ core.Actor = (*MyObject)(nil)

func NewMyObject(assets *asset.Asset) *MyObject {
	return &MyObject{
		assets: assets,
	}
}

func (o *MyObject) Update() error {
	o.ticks++
	return nil
}

func (o *MyObject) Draw(screen *ebiten.Image) {
	textFace := o.assets.TextFace()

	screen.Fill(color.RGBA{R: 22, G: 26, B: 34, A: 255})
	drawText(screen, "これは日本語です "+string(mdi.Home), textFace, 96, 140, color.RGBA{R: 240, G: 244, B: 255, A: 255})
	drawText(screen, "NotoSansMonoCJKjp-Bold + Material Design Icons fallback", textFace, 96, 230, color.RGBA{R: 168, G: 176, B: 202, A: 255})
	drawText(screen, o.assets.FontSource(), textFace, 96, 300, color.RGBA{R: 126, G: 138, B: 170, A: 255})

	o.drawIconImage(screen, 160, 500, 1.4, color.RGBA{R: 110, G: 231, B: 183, A: 255})
}

func (o *MyObject) Alive() bool {
	return true
}

func (o *MyObject) drawIconImage(screen *ebiten.Image, x, y, scale float64, tint color.RGBA) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleWithColor(tint)
	screen.DrawImage(o.assets.IconImage(mdi.Home), op)
}

func drawText(dst *ebiten.Image, s string, face text.Face, x, y int, c color.Color) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	op.ColorScale.ScaleWithColor(c)
	text.Draw(dst, s, face, op)
}
