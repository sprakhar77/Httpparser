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
		log.Println("Maximum concurrent operations allowed is 100. Using 100 workers")
		concurrency = 100
	}
	if concurrency <= 0 {
		log.Println("Minimum concurrent can not be negetive or 0, setting it to 1")
		concurrency = 1
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
