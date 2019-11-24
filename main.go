package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"

	"github.com/srleyva/simple-bus/pkg/message"
)

func main() {
	log.Print("simple-bus starting up")
	notifyFunc := func(message *message.Message) {
		log.Print(message.GetMessage())
	}
	recievers := []*message.BusNodeReciever{}

	for i := 0; i < 100; i++ {
		name := fmt.Sprintf("%d-recv", i)
		log.Printf("Initializing reciever: %s", name)
		recievers = append(recievers, message.NewBusNodeReciever(name, notifyFunc))
	}

	log.Print("initilizing node bus")
	bus, err := message.NewBusNode(recievers)
	if err != nil {
		log.Fatalf("error initializing bus: %s", err)
	}

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go bus.Update(sigs, done)
	go func() {
		mux := mux.NewRouter()
		AttachProfiler(mux)
		mux.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
			var message message.Message
			if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			bus.MessageBus.SendMessage(&message)
		})
		http.ListenAndServe(":6060", mux)
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			message := message.NewMessage(fmt.Sprintf("event %d!", i))
			bus.MessageBus.SendMessage(message)
		}
	}()
	<-done
	log.Print("Shutting Down")

}

func AttachProfiler(router *mux.Router) {
	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)

	// Manually add support for paths linked to by index page at /debug/pprof/
	router.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	router.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	router.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	router.Handle("/debug/pprof/block", pprof.Handler("block"))
}
