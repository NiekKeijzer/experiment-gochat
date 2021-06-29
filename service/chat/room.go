package chat

import "log"

type Room struct {
	clients map[string]*Client

	Name       string
	Broadcast  chan *OutMessage
	Register   chan *Client
	Unregister chan *Client
}

func (r *Room) run() {
	for {
		log.Println("asdf")
		select {
		case client := <-r.Register:
			log.Println("registering client")
			r.clients[client.Name] = client
			client.Room = r
		case client := <-r.Unregister:
			log.Println("unregistering client")
			if _, ok := r.clients[client.Name]; ok {
				delete(r.clients, client.Name)
			}
		case message := <-r.Broadcast:
			log.Println("Broadcasting message")
			for name, client := range r.clients {
				select {
				case client.Send <- message:
				default:
					delete(r.clients, name)
				}
			}
		}
	}
}
