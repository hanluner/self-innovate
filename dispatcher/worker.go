package main

import (
	"fmt"
	"time"
)

// Worker - worker struct
type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit chan bool
}

// NewWorker - create worker instance.
func NewWorker(workerPool chan chan Job) *Worker {
	return &Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit: make(chan bool),
	}
}

func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel
			select {
			case job := <-w.JobChannel:
				time.Sleep(500 * time.Millisecond)
				fmt.Println("request complete: ", job)
			case <-w.quit:
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}