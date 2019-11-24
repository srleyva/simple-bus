package message

import (
	"fmt"
	"log"
	"os"
	"time"
)

type BusNodeReciever struct {
	name     string
	messages []*Message
	onNotify func(*Message)
}

func NewBusNodeReciever(name string, notifyFunc func(*Message)) *BusNodeReciever {
	return &BusNodeReciever{name, []*Message{}, notifyFunc}
}

func (r *BusNodeReciever) Notify(message *Message) {
	log.Printf("MESSAGE FROM RECV %s", r.name)
	if message == nil {
		log.Print("message is nil")
		return
	}
	r.onNotify(message)
}

type BusNode struct {
	MessageBus *MessageBus
}

func NewBusNode(recvs []*BusNodeReciever) (*BusNode, error) {
	if len(recvs) < 1 {
		return nil, fmt.Errorf("Not enough recvs: %d", len(recvs))
	}

	bus := &MessageBus{}
	for _, recv := range recvs {
		bus.AddReceiver(recv)
	}

	return &BusNode{bus}, nil
}

func (b *BusNode) Update(sigs chan os.Signal, done chan bool) {
	log.Print("bus is starting")
	for true {
		select {
		case sig := <-sigs:
			log.Printf("sig recv: %s", sig.String())
			done <- true
		default:
			b.MessageBus.Notify()
			time.Sleep(1 * time.Second)
		}
	}
}
