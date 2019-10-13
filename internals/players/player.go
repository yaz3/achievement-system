package players

// Player is used as the main player struct
type Player struct {
	ID           int64  `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	Wins         int64  `json:"wins" db:"wins"`
	Losses       int64  `json:"losses" db:"losses"`
	Time_played  int64  `json:"time_played" db:"time_played"` // in minutes
	Achievements string `json:"achievements" db:"achievements"`
}
