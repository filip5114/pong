package models

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	InitPaddleHeight = 100.0
	InitPaddleWidth  = 10.0
)

type Paddle struct {
	Position
	H, W    float32
	Yv      float32
	Color   color.Color
	Img     *ebiten.Image
	pressed keysPressed
	Up      ebiten.Key
	Down    ebiten.Key
	Score   int
}

type keysPressed struct {
	up   bool
	down bool
}

func (p *Paddle) Update() {
	if inpututil.IsKeyJustPressed(p.Up) {
		p.pressed.down = false
		p.pressed.up = true
	} else if inpututil.IsKeyJustReleased(p.Up) {
		p.pressed.up = false
	}
	if inpututil.IsKeyJustPressed(p.Down) {
		p.pressed.up = false
		p.pressed.down = true
	} else if inpututil.IsKeyJustReleased(p.Down) {
		p.pressed.down = false
	}

	if p.pressed.up {
		p.Y -= p.Yv
	} else if p.pressed.down {
		p.Y += p.Yv
	}

	if p.Y < 0 {
		p.Y = 0
	} else if p.Y+p.H > 480.0 {
		p.Y = 480.0 - p.H
	}
}

func (p *Paddle) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(p.X), float64(p.Y))
	p.Img.Fill(p.Color)
	screen.DrawImage(p.Img, options)
}
