package main

import "sync"

type Task struct {
	f   func() error
	err error
}

func (t *Task) Run(wg *sync.WaitGroup) {
	t.err = t.f()
	wg.Done()
}

func NewTask(f func() error) *Task {
	return &Task{f: f}
}
