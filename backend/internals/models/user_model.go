package models

import "time"

type GameStatus int

const (
	IN_PROGRESS GameStatus = iota
	COMPLETED
	ABANDONED
	TIME_UP
	PLAYER_EXIT
)

type GameResult int

const (
	WHITE_WINS GameResult = iota
	BLACK_WINS
	DRAW
)

type TimeControl int

const (
	CLASSICAL TimeControl = iota
	RAPID
	BLITZ
	BULLET
)

type User struct {
	ID           string
	Username     string
	Name         string
	Email        string
	Password     string
	Rating       int
	GamesAsWhite []Game
	GamesAsBlack []Game
	CreatedAt    time.Time
	LastLogin    time.Time
}

type Game struct {
	ID            string
	WhitePlayerID string
	BlackPlayerID string
	WhitePlayer   User
	BlackPlayer   User
	Status        GameStatus
	Result        GameResult
	TimeControl   TimeControl
	StartingFEN   string
	CurrentFEN    string
	StartAt       time.Time
	EndAt         time.Time
	Moves         []Move
	Opening       string
}

type Move struct {
	ID         string
	GameID     string
	Game       Game
	MoveNumber int
	From       string
	To         string
	Comments   string
	Before     string
	After      string
	TimeTaken  int
	CreatedAt  time.Time
	SAN        string
}
