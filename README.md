# Pong
Simple Pong game written in Go using Ebiten game library.

The project was created to learn and get used to Go in a fun way.

Features:
* two players
* resizeable screen
* ball velocity incremental increase

# Run
Clone the repo and run commands:
`go build`
`./pong`

# Settings
All settings are defined in main.go file:
* initBallVelocity - initial ball velocity
* initBallStep - ball velocity increase step
* initMaxScore - score to win
* Up - key to move paddle up, in init() function for player1 and player2
* Down - key to move paddle down, in init() function for player1 and player2


