package message

import (
	"testing"
)

func TestMessage(t *testing.T) {
	t.Run("test Setting and retrieving of message", func(t *testing.T) {
		event := "New Event"
		message := NewMessage(event)

		if message.GetMessage() != event {
			t.Errorf("message is incorrect! Expected: %s, Actual: %s", event, message.GetMessage())
		}
	})
}
