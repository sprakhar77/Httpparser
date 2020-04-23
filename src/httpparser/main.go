package main

import (
	"flag"
)

func main() {

	var concurrency int
	flag.IntVar(&concurrency, "concurrency", 10, "Number of concurrency tasks allowed")
	flag.Parse()

	tasks := make([]*Task, 0)
	for index := range flag.Args() {
		url := flag.Args()[index]
		tasks = append(tasks, NewTask(func() error { return Process(url) }))
	}

	workerPool := NewWorkerPool(tasks, concurrency)
	workerPool.Run()
}
