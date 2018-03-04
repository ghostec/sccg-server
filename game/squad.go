package game

import "github.com/ghostec/sccg-server/math"

// SquadPlayer struct
type SquadPlayer struct {
	Position math.Vec2 `json:"position"`
}

// Squad struct
type Squad struct {
	Players []SquadPlayer `json:"players"`
}

// NewSquad ctor
func NewSquad() *Squad {
	return &Squad{
		Players: []SquadPlayer{
			SquadPlayer{
				Position: math.Vec2{0.0, 0.0},
			},
		},
	}
}
