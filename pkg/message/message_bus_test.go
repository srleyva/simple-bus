package message

import (
	"testing"
)

type MockRecv struct {
	messages []*Message
}

func (m *MockRecv) Notify(message *Message) {
	m.messages = append(m.messages, message)
}

func (m *MockRecv) GetMessages() []*Message {
	return m.messages
}

func TestMessageBus(t *testing.T) {
	t.Run("test adding recievers", func(t *testing.T) {
		mockRecv1 := &MockRecv{[]*Message{}}
		mockRecv2 := &MockRecv{[]*Message{}}
		bus := &MessageBus{}
		bus.AddReceiver(mockRecv1)
		bus.AddReceiver(mockRecv2)
		if len(bus.GetReceivers()) != 2 {
			t.Errorf("wrong number of receivers! Expected: %d, Actual: %d", 2, bus.GetReceivers())
		}
	})
	t.Run("test writing to recievers", func(t *testing.T) {
		mockRecv1 := &MockRecv{[]*Message{}}
		mockRecv2 := &MockRecv{[]*Message{}}
		bus := &MessageBus{}
		bus.AddReceiver(mockRecv1)
		bus.AddReceiver(mockRecv2)
		message1 := &Message{"event 1"}
		message2 := &Message{"event 2"}
		bus.SendMessage(message1)
		bus.SendMessage(message2)
		bus.Notify()
		if mockRecv1.GetMessages()[0] != message1 || mockRecv2.GetMessages()[0] != message2 {
			t.Errorf("messages not received!")
		}

	})
}
