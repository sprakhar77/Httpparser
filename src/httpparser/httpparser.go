package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
)

func ComputeMD5Hash(data string) string {
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

func SendRequest(url string) error {
	response, error := http.Get(url)

	if error != nil {
		return error
	}
	defer response.Body.Close()
	body, error := ioutil.ReadAll(response.Body)

	if error != nil {
		return error
	}
	log.Println(url, ComputeMD5Hash(string(body)))
	return nil
}

func main() {

	var concurrency int
	flag.IntVar(&concurrency, "concurrency", 10, "Number of concurrency tasks allowed")
	flag.Parse()

	tasks := make([]*Task, 0)
	for _, url := range flag.Args() {
		tasks = append(tasks, NewTask(func() error { return SendRequest(url) }))
	}

	workerPool := NewWorkerPool(tasks, concurrency)
	workerPool.Run()

	for _, task := range workerPool.tasks {
		if task.err != nil {
			log.Fatal(task.err)
		}
	}
}
