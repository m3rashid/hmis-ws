package main

import (
	"github.com/m3rashid/hmis-ws/router"
	"github.com/m3rashid/hmis-ws/ws"
)

func main() {
	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(wsHandler)
	router.Start("0.0.0.0:4000")
}
