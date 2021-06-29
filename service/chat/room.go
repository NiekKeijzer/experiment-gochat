package chat

type Room struct {
	clients map[string]*Client

	Name       string
	Broadcast  chan *OutMessage
	Register   chan *Client
	Unregister chan *Client
}

func (r *Room) run() {
	for {
		select {
		case client := <-r.Register:
			r.clients[client.Name] = client
			client.Room = r
		case client := <-r.Unregister:
			if _, ok := r.clients[client.Name]; ok {
				delete(r.clients, client.Name)
			}
		case message := <-r.Broadcast:
			for name, client := range r.clients {
				if message.SenderId == client.id {
					continue
				}

				select {
				case client.Send <- message:
				default:
					delete(r.clients, name)
				}
			}
		}
	}
}
