package chat

type InMessage struct {
	Message string
}

type OutMessage struct {
	Message  string `json:"message"`
	Sender   string `json:"sender"`
	SenderId string `json:"-"`
}
