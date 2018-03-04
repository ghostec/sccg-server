package api

import (
	"context"
	"net"
	"net/http"

	"github.com/gorilla/websocket"
)

// App struct
type App struct {
	Address  string
	Server   *http.Server
	Upgrader websocket.Upgrader
}

// NewApp ctor
func NewApp(a string) (*App, error) {
	app := &App{
		Address: a,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}

	err := app.configureApp()
	if err != nil {
		return nil, err
	}
	return app, nil
}

func (a *App) configureApp() error {
	err := a.configureServer()
	if err != nil {
		return err
	}
	return nil
}

func (a *App) configureServer() error {
	sm := http.NewServeMux()
	sm.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		a.WsHandler(w, r)
	})
	a.Server = &http.Server{Addr: a.Address, Handler: sm}
	return nil
}

// ListenAndServe method
func (a *App) ListenAndServe() error {
	listener, err := net.Listen("tcp", a.Address)
	if err != nil {
		return err
	}
	err = a.Server.Serve(listener)
	if err != nil {
		listener.Close()
		return err
	}
	return nil
}

// Stop method
func (a *App) Stop() {
	a.Server.Shutdown(context.Background())
}
