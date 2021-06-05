package main

type Dispatcher struct {
	WorkerPool chan chan Job
	Max int
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	return &Dispatcher{
		WorkerPool: make(chan chan Job, maxWorkers),
		Max: maxWorkers,
	}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.Max; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <- JobQueue: // 应该有更好的JobQueue初始化方式
			go func(job Job) {
				jobChannel := <- d.WorkerPool // try to get a worker_pool
				jobChannel <- job
			}(job)
		}
	}
}