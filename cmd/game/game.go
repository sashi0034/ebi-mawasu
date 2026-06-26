package main

import (
	"bytes"
	"fmt"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	renderWidth  = 1920
	renderHeight = 1080
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

func renderIconImage(r rune, face text.Face, size int) *ebiten.Image {
	img := ebiten.NewImage(size, size)
	drawText(img, string(r), face, 32, 32, color.White)
	return img
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
