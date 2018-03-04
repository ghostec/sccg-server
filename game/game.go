package game

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/ghostec/sccg-server/math"
)

// Game struct
type Game struct {
	FPS   float64     `json:"-"`
	Mutex *sync.Mutex `json:"-"`
	Squad *Squad      `json:"squad"`
}

// NewGame ctor
func NewGame() *Game {
	return &Game{
		FPS:   30.0,
		Mutex: &sync.Mutex{},
		Squad: NewSquad(),
	}
}

// Update method
func (g *Game) Update() error {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()

	g.Squad.Update()

	return nil
}

// Loop method
func (g *Game) Loop() {
	for {
		err := math.WithFrameInterval(g.FPS, func() error {
			return g.Update()
		})
		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}
}

// Snapshot func
func (g Game) Snapshot() ([]byte, error) {
	g.Mutex.Lock()
	b, err := json.Marshal(g)
	g.Mutex.Unlock()

	if err != nil {
		return []byte{}, err
	}
	return b, nil
}
