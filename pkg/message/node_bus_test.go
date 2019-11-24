package message

import (
	"fmt"
	"testing"
)

func TestNodeBus(t *testing.T) {
	t.Run("test creating a new node bus", func(t *testing.T) {
		recievers := []*BusNodeReciever{
			&BusNodeReciever{[]*Message{}, nil},
			&BusNodeReciever{[]*Message{}, nil},
			&BusNodeReciever{[]*Message{}, nil},
			&BusNodeReciever{[]*Message{}, nil},
			&BusNodeReciever{[]*Message{}, nil},
		}

		busNode, err := NewBusNode(recievers, func(message *Message) {
			fmt.Printf("MESSAGE: %s\n", message.GetMessage())
		})
		if err != nil {
			t.Errorf("error where not expected! %s", err)
		}

		if len(busNode.MessageBus.GetReceivers()) != len(recievers) {
			t.Errorf("Incorrect number of receivers: expected: %d actual: %d", len(recievers), len(busNode.MessageBus.GetReceivers()))
		}
	})
}
