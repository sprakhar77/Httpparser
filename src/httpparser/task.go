package main

import "sync"

type Task struct {
	job func() error
	err error
}

func (task *Task) Run(waitGroup *sync.WaitGroup) {
	task.err = task.job()
	waitGroup.Done()
}

func NewTask(job func() error) *Task {
	return &Task{job: job}
}
