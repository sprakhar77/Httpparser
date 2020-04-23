package main

import (
	"log"
	"sync"
)

type WorkerPool struct {
	tasks       []*Task
	concurrency int
	tasksChan   chan *Task
	waitGroup   sync.WaitGroup
}

func NewWorkerPool(tasks []*Task, concurrency int) *WorkerPool {
	if concurrency > 100 {
		log.Println("Concurrency is limited to 100.")
		concurrency = 100
	}
	return &WorkerPool{
		tasks:       tasks,
		concurrency: concurrency,
		tasksChan:   make(chan *Task),
	}
}

func (workerPool *WorkerPool) Run() {
	for i := 0; i < workerPool.concurrency; i++ {
		go workerPool.work()
	}
	workerPool.waitGroup.Add(len(workerPool.tasks))
	for _, task := range workerPool.tasks {
		workerPool.tasksChan <- task
	}
	close(workerPool.tasksChan)
	workerPool.waitGroup.Wait()
}

func (workerPool *WorkerPool) work() {
	for task := range workerPool.tasksChan {
		task.Run(&workerPool.waitGroup)
	}
}
