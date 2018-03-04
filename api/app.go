package api

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"

	"github.com/gorilla/websocket"
)

// App struct
type App struct {
	Address string
	Server  *http.Server
}

// NewApp ctor
func NewApp(a string) (*App, error) {
	app := &App{
		Address: a,
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

var addr = flag.String("addr", "localhost:8080", "http service address")
var upgrader = websocket.Upgrader{}

func main() {
	a, err := NewApp("localhost:8080")
	if err != nil {
		fmt.Println(err.Error())
	}
	a.ListenAndServe()
}
