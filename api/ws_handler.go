package api

import (
	"fmt"
	"net/http"

	"github.com/ghostec/sccg-server/game"
	"github.com/ghostec/sccg-server/math"
	"github.com/gorilla/websocket"
)

// WsHandler func
func (a *App) WsHandler(w http.ResponseWriter, r *http.Request) {
	c, err := a.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer c.Close()

	g := game.NewGame()
	go g.Loop()

	for {
		err := math.WithFrameInterval(g.FPS, func() error {
			b, err := g.Snapshot()
			if err != nil {
				return err
			}
			err = c.WriteMessage(websocket.BinaryMessage, b)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}
}
