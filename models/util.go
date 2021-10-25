package models

type Position struct {
	X, Y float32
}

type GameState byte

const (
	StartState GameState = iota
	PlayState
	GameOverState
)
