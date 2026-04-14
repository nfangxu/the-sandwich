package ws

import (
	"log"

	socketio "github.com/googollee/go-socket.io"
)

type Hub struct {
	Server *socketio.Server
}

func NewHub() (*Hub, error) {
	server := socketio.NewServer(nil)

	hub := &Hub{Server: server}

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Printf("Socket connected: %s", s.ID())
		return nil
	})

	server.OnEvent("/", "message", func(s socketio.Conn, msg string) {
		log.Printf("Received message from %s: %s", s.ID(), msg)
		s.Emit("message", msg)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Printf("Socket disconnected: %s, reason: %s", s.ID(), reason)
	})

	return hub, nil
}

func (h *Hub) Run() {
	h.Server.Serve()
}

func (h *Hub) Broadcast(event string, data interface{}) {
	h.Server.BroadcastToRoom("/", "game", event, data)
}
