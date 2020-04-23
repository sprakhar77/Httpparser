package main

import "testing"

func TestMaximumNumberOfConcurentTasks(t *testing.T) {
	tasks := make([]*Task, 0)
	tasks = append(tasks, NewTask(func() error { return nil }))
	workerPool := NewWorkerPool(tasks, 200)
	if workerPool.concurrency != 100 {
		t.Errorf("Maximum number of tasks is limited to 100, got: %d", workerPool.concurrency)
	}
}

func TestMinimumNumberOfConcurentTasks(t *testing.T) {
	tasks := make([]*Task, 0)
	tasks = append(tasks, NewTask(func() error { return nil }))
	workerPool := NewWorkerPool(tasks, -2)
	if workerPool.concurrency != 1 {
		t.Errorf("Minimum number of tasks should be 0, got: %d", workerPool.concurrency)
	}
}
