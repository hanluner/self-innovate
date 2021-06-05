package main

import (
	"log"
	"net/http"
	"time"
)

const (
	MaxWorker = 10
	MaxQueue = 20
)

var JobQueue chan Job

func init() {
	JobQueue = make(chan Job, MaxQueue)
}

func main() {
	dispatcher := NewDispatcher(MaxWorker)
	dispatcher.Run()

	http.HandleFunc("/payload", payloadHandler)
	log.Fatal(http.ListenAndServe(":8899", nil))
}

func payloadHandler(w http.ResponseWriter, r *http.Request) {
	job := Job{Pl: Payload{Name: "taskName"}}
	JobQueue <- job
	time.Sleep(2000 * time.Millisecond)
	w.Write([]byte("success"))
}
