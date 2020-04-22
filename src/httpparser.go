package main

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
	tasks := make([]*Task, 0)
	for _, url := range os.Args[1:] {
		tasks = append(tasks, NewTask(func() error { return SendRequest(url) }))
	}

	pool := NewPool(tasks, 10)
	pool.Run()

	for _, task := range pool.tasks {
		if task.err != nil {
			log.Fatal(task.err)
		}
	}
}