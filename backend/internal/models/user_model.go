package models

import "time"

type GameStatus string
type GameResult string
type TimeControl string

const (
	IN_PROGRESS GameStatus = "in_progress"
	COMPLETED   GameStatus = "completed"
	ABANDONED   GameStatus = "abandoned"
	TIME_UP     GameStatus = "time_up"
	PLAYER_EXIT GameStatus = "player_exit"
)

const (
	WHITE_WINS GameResult = "white_wins"
	BLACK_WINS GameResult = "black_wins"
	DRAW       GameResult = "draw"
)

const (
	CLASSICAL TimeControl = "classical"
	RAPID     TimeControl = "rapid"
	BLITZ     TimeControl = "blitz"
	BULLET    TimeControl = "bullet"
)

type User struct {
	ID           string    `json:"id"`
	Username     string    `json:"username"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     string    `json:"-"`
	Rating       int       `json:"rating"`
	GamesAsWhite []Game    `json:"games_as_white,omitempty"`
	GamesAsBlack []Game    `json:"games_as_black,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	LastLogin    time.Time `json:"last_login"`
}

type Game struct {
	ID            string      `json:"id"`
	WhitePlayerID string      `json:"white_player_id"`
	BlackPlayerID string      `json:"black_player_id"`
	WhitePlayer   *User       `json:"white_player,omitempty"`
	BlackPlayer   *User       `json:"black_player,omitempty"`
	Status        GameStatus  `json:"status"`
	Result        *GameResult `json:"result,omitempty"`
	TimeControl   TimeControl `json:"time_control"`
	StartingFEN   string      `json:"starting_fen"`
	CurrentFEN    string      `json:"current_fen"`
	StartAt       time.Time   `json:"start_at"`
	EndAt         *time.Time  `json:"end_at,omitempty"`
	Moves         []Move      `json:"moves,omitempty"`
	Opening       string      `json:"opening"`
}

type Move struct {
	ID         string    `json:"id"`
	GameID     string    `json:"game_id"`
	MoveNumber int       `json:"move_number"`
	From       string    `json:"from"`
	To         string    `json:"to"`
	SAN        string    `json:"san"`
	Before     string    `json:"before"`
	After      string    `json:"after"`
	Comments   string    `json:"comments"`
	TimeTaken  int       `json:"time_taken"`
	CreatedAt  time.Time `json:"created_at"`
}
