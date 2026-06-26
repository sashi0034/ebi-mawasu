package asset

import (
	"bytes"
	"fmt"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Asset struct {
	textFace   text.Face
	iconFace   text.Face
	fontSource string
	iconImages map[iconImageKey]*ebiten.Image
}

type iconImageKey struct {
	mdiCode rune
	size    int
}

func NewAsset() (*Asset, error) {
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

	return &Asset{
		textFace:   textFace,
		iconFace:   iconFace,
		fontSource: fontSource,
		iconImages: make(map[iconImageKey]*ebiten.Image),
	}, nil
}

func (a *Asset) TextFace() text.Face {
	return a.textFace
}

func (a *Asset) FontSource() string {
	return a.fontSource
}

func (a *Asset) IconImage(mdiCode rune) *ebiten.Image {
	size := 256
	key := iconImageKey{
		mdiCode: mdiCode,
		size:    size,
	}
	if image := a.iconImages[key]; image != nil {
		return image
	}

	if a.iconImages == nil {
		a.iconImages = make(map[iconImageKey]*ebiten.Image)
	}

	image := renderIconImage(mdiCode, a.iconFace, size)
	a.iconImages[key] = image
	return image
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

func renderIconImage(r rune, face text.Face, size int) *ebiten.Image {
	image := ebiten.NewImage(size, size)

	op := &text.DrawOptions{}
	op.PrimaryAlign = text.AlignCenter
	op.SecondaryAlign = text.AlignCenter
	op.GeoM.Translate(float64(size)/2, float64(size)/2)
	op.ColorScale.ScaleWithColor(color.White)

	text.Draw(image, string(r), face, op)

	return image
}
