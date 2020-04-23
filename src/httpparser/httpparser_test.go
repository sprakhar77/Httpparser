package main

import "testing"

func TestMaxConcurentTasks(t *testing.T) {
	tasks := make([]*Task, 0)
	tasks = append(tasks, NewTask(func() error { return nil }))
	workerPool := NewWorkerPool(tasks, 200)
	if workerPool.concurrency != 100 {
		t.Errorf("Maximum number of tasks was incorrect, got: %d, want: %d", workerPool.concurrency, 100)
	}
}
