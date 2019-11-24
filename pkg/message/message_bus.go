package message

import "log"

type receiver interface {
	Notify(*Message)
}

type MessageBus struct {
	messages  []*Message
	receivers []receiver
}

func (m *MessageBus) GetReceivers() []receiver {
	return m.receivers
}

func (m *MessageBus) AddReceiver(recv receiver) {
	m.receivers = append(m.receivers, recv)
}

func (m *MessageBus) SendMessage(message *Message) {
	m.messages = append(m.messages, message)
}

func (m *MessageBus) Notify() {
	for len(m.messages) != 0 {
		for _, recv := range m.receivers {
			if len(m.messages) != 0 {
				message := m.messages[0]
				m.messages = m.messages[1:]
				recv.Notify(message)
			} else {
				log.Print("No messages!")
			}
		}
	}
}
