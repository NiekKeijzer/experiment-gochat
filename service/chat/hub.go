package chat

import (
	"github.com/gorilla/websocket"
)

const (
	defaultRoomName = "default"
	defaultUserName = "anonymous"
)

type Hub struct {
	rooms map[string]*Room

	Register   chan *Client
	Unregister chan *Client
}

func NewHub() *Hub {
	hub := &Hub{
		rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}

	defaultRoom := &Room{
		clients:    make(map[string]*Client),
		Name:       defaultRoomName,
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
	go defaultRoom.run()

	hub.rooms[defaultRoom.Name] = defaultRoom

	return hub
}

func (h *Hub) NewClient(conn *websocket.Conn) *Client {
	client := &Client{
		Conn: conn,
		Name: defaultRoomName,
		Hub:  h,
		Send: make(chan *OutMessage),
	}

	h.Register <- client

	return client
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			if defaultRoom, ok := h.rooms[defaultRoomName]; ok {
				defaultRoom.Register <- client
			}
		case client := <-h.Unregister:
			client.Room.Unregister <- client

			close(client.Send)
		}
	}
}
