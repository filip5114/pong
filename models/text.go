package models

import (
	"fmt"
	"image/color"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

const (
	fontSize  = 30
	smallSzie = 10
)

var (
	Font1 font.Face
)

func init() {
	tt, err := truetype.Parse(fonts.ArcadeN_ttf)
	if err != nil {
		log.Fatal(err)
	}
	var dpi float64 = 48
	Font1 = truetype.NewFace(tt, &truetype.Options{
		Size:    float64(fontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

func drawScore(w int, h int, score1 int, score2 int, color color.Color, screen *ebiten.Image) {
	text.Draw(screen, fmt.Sprint(score1), Font1, w/20, h/10, color)
	text.Draw(screen, fmt.Sprint(score2), Font1, (w/20)*18, h/10, color)
}

func Draw(state GameState, player1 *Paddle, player2 *Paddle, color color.Color, screen *ebiten.Image) {
	var msg string
	var x, y int
	w, h := screen.Size()
	x = w / 2
	font := Font1
	drawScore(w, h, player1.Score, player2.Score, color, screen)
	switch state {
	case StartState:
		y = (h / 4) * 3
		msg = "Press space to start"
	case PlayState:
		x = 0
		y = 0
		msg = ""
	case InterState:
		y = h / 4
		if player1.LastPoint {
			msg = "Score for Player 1 !"
		} else if player2.LastPoint {
			msg = "Score for Player 2 !"
		}
	case GameOverState:
		y = h / 4
		if player1.Winner {
			msg = "! Player 1 Won !"
		} else if player2.Winner {
			msg = "! Player 2 Won !"
		}
	}
	textSize := text.BoundString(Font1, msg).Dx()
	text.Draw(screen, msg, font, x-textSize/2, y, color)

}
