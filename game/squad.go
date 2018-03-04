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
				Position: math.Vec2{X: 0.0, Y: 0.0},
			},
		},
	}
}

// Update method for Squad
func (s *Squad) Update() {
	s.Players[0].Update()
}

// Update method for SquadPlayer
func (p *SquadPlayer) Update() {
	p.Position.X++
}
