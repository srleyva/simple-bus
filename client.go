package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	i, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	for j := 0; j < i; j++ {
		log.Printf("request %d", j)
		wg.Add(1)
		NewReq(j, &wg)
	}
	wg.Wait()
}

func NewReq(current int, wg *sync.WaitGroup) {
	body, err := json.Marshal(map[string]string{
		"event": fmt.Sprintf("request_%d", current),
	})
	if err != nil {
		log.Print(err)
	}
	resp, err := http.Post("http://localhost:6060/message", "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Print(err)
		wg.Done()
		return
	}

	if resp.StatusCode != 200 {
		log.Printf("Non 200 returned! %d", resp.StatusCode)
	}
	wg.Done()
}
