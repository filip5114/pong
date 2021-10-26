package models

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Ball struct {
	Position
	Radius float32
	Color  color.Color
	Img    *ebiten.Image
	Xv, Yv float32
	Step   float32
}

const (
	InitBallRadius = 5.0
)

func (b *Ball) Update(w int, h int, leftPaddle *Paddle, rightPaddle *Paddle) {
	b.X += b.Xv
	b.Y += b.Yv

	if b.Y < 0 {
		b.Yv = -b.Yv
		b.Y = 0
	} else if b.Y+b.Radius*2 > float32(h) {
		b.Yv = -b.Yv
		b.Y = float32(h) - b.Radius*2
	}

	if b.X < leftPaddle.X+b.Radius*2 && b.X > leftPaddle.X-b.Radius*2 && b.Y > leftPaddle.Y && b.Y < leftPaddle.Y+leftPaddle.H {
		b.Xv = -b.Xv + b.Step
		b.X = leftPaddle.X + b.Radius*2
	} else if b.X > rightPaddle.X-b.Radius*2 && b.X < rightPaddle.X+b.Radius*2 && b.Y > rightPaddle.Y && b.Y < rightPaddle.Y+leftPaddle.H {
		b.Xv = -b.Xv - b.Step
		b.X = rightPaddle.X - b.Radius*2
	}

	if math.Abs(float64(b.Yv))/math.Abs(float64(b.Xv)) < 0.8 {
		if b.Yv < 0 {
			b.Yv -= b.Step * 2
		} else {
			b.Yv += b.Step * 2
		}
	}
}

func (b *Ball) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(b.X), float64(b.Y))
	b.Img.Fill(b.Color)
	screen.DrawImage(b.Img, options)
}
