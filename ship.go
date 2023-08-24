package main

import (
	"image"
	"image/color"
	"image/draw"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	_ "golang.org/x/image/bmp"
)

type Ship struct {
	image  *ebiten.Image
	width  int
	height int
	x      float64
	y      float64
}

func NewShip(screenWidth, screenHeight int, cfg *Config) *Ship {
	rect := image.Rect(0, 0, 8, 8)
	img := ebiten.NewImageWithOptions(rect, nil)
	img.Fill(cfg.BulletColor)

	width, height := img.Size()
	ship := &Ship{
		image:  img,
		width:  width,
		height: height,
		x:      float64(screenWidth-width) / 2,
		y:      float64(screenHeight - height),
	}

	return ship
}

func (ship *Ship) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(ship.x, ship.y)
	screen.DrawImage(ship.image, op)
}
