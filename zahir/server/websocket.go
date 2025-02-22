package server

import (
	"log"
	"net/http"
	"time"
	"zahir/player"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var lastStep = -1

// /v1/ws
func wsState(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		state := player.GetCurrentPlayerState()
		if true {
			lastStep = state.Step
			err := conn.WriteJSON(state)
			if err != nil {
				log.Println("WebSocket write error:", err)
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
}
