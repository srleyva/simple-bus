package message

type Message struct {
	MessageEvent string `json:"event"`
}

func NewMessage(event string) *Message {
	return &Message{event}
}

func (m *Message) GetMessage() string {
	return m.MessageEvent
}
