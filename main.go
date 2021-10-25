package main

import (
	"image/color"
	"log"

	"github.com/filip5114/pong/models"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth          = 640
	screenHeight         = 480
	initBallPositionX    = float32(screenWidth / 2)
	initBallPositionY    = float32(screenHeight / 2)
	initPlayer1PositionX = float32(screenWidth / 4)
	initPlayer1PositionY = float32((screenHeight / 2) - (models.InitPaddleHeight / 2))
	initPlayer2PositionX = float32((screenWidth / 4) * 3)
	initPlayer2PositionY = float32((screenHeight / 2) - (models.InitPaddleHeight / 2))
	initBallVelocity     = 3.0
)

type Game struct {
	ball    *models.Ball
	player1 *models.Paddle
	player2 *models.Paddle
	state   models.GameState
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.ball.Draw(screen)
	g.player1.Draw(screen)
	g.player2.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Update() error {
	switch g.state {
	case models.StartState:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.state = models.PlayState
		}
	case models.PlayState:
		g.player1.Update()
		g.player2.Update()
		g.ball.Update(g.player1, g.player2)

		if g.ball.X < 0 {
			g.player2.Score++
			g.reset()
		} else if g.ball.X > screenWidth {
			g.player1.Score++
			g.reset()
		}
	}
	return nil
}

func (g *Game) reset() {
	g.ball.Position = models.Position{
		X: initBallPositionX,
		Y: initBallPositionY,
	}
	g.ball.Xv = initBallVelocity
	g.ball.Yv = initBallVelocity
	g.player1.Position = models.Position{
		X: initPlayer1PositionX,
		Y: initPlayer1PositionY,
	}
	g.player2.Position = models.Position{
		X: initPlayer2PositionX,
		Y: initPlayer2PositionY,
	}
	g.state = models.StartState
}

func (g *Game) init() {
	g.state = models.StartState
	g.ball = &models.Ball{
		Position: models.Position{
			X: initBallPositionX,
			Y: initBallPositionY},
		Xv:     initBallVelocity,
		Yv:     initBallVelocity,
		Radius: models.InitBallRadius,
		Color:  color.RGBA{255, 255, 255, 255},
	}
	g.player1 = &models.Paddle{
		Position: models.Position{
			X: initPlayer1PositionX,
			Y: initPlayer1PositionY},
		Yv:    5.0,
		H:     models.InitPaddleHeight,
		W:     models.InitPaddleWidth,
		Color: color.RGBA{255, 255, 255, 255},
		Up:    ebiten.KeyUp,
		Down:  ebiten.KeyDown,
		Score: 0,
	}
	g.player2 = &models.Paddle{
		Position: models.Position{
			X: initPlayer2PositionX,
			Y: initPlayer2PositionY},
		Yv:    5.0,
		H:     models.InitPaddleHeight,
		W:     models.InitPaddleWidth,
		Color: color.RGBA{255, 255, 255, 255},
		Up:    ebiten.KeyA,
		Down:  ebiten.KeyZ,
		Score: 0,
	}
	g.ball.Img = ebiten.NewImage(int(g.ball.Radius)*2, int(g.ball.Radius)*2)
	g.player1.Img = ebiten.NewImage(int(g.player1.W), int(g.player1.H))
	g.player2.Img = ebiten.NewImage(int(g.player2.W), int(g.player2.H))
}

func main() {
	g := &Game{}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Pong")
	g.init()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
