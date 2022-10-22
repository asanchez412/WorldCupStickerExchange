package domain

import "time"

//Sticker represents the main features of a world cup sticker album
type Sticker struct {
	Id          int       `json:"id"`
	TeamSticker bool      `json:"team_sticker"`
	Country     string    `json:"country"`
	Height      float32   `json:"height"`
	Weight      int       `json:"weight"`
	Position    uint      `json:"position"`
	DebutYear   string    `json:"debut_year"`
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

const (
	Attacker   = 3
	Midfielder = 2
	Defender   = 1
	Goalkeeper = 0
)

//validate validates whether or not the sticker values are correct
func (s *Sticker) validate() bool {
	if s.Position > 3 {
		return false
	}
	if s.Height < 1 {
		return false
	}
	if s.Weight < 1 {
		return false
	}
	return true
}
