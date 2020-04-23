package main

import "testing"

func TestComputeMD5HashWithEmptyString(t *testing.T) {
	result := ComputeMD5Hash("")
	if result == "" || result != "d41d8cd98f00b204e9800998ecf8427e" {
		t.Errorf("MD5Hash for empty string shoud be equal to d41d8cd98f00b204e9800998ecf8427e")
	}
}

func TestComputeMD5HashIsConsistentForSameString(t *testing.T) {
	firstResult := ComputeMD5Hash("some random string")
	secondResult := ComputeMD5Hash("some random string")
	if firstResult != secondResult {
		t.Errorf("MD5Hash should be same for the same input, found different first: %s, second: %s", firstResult, secondResult)
	}
}

func TestSendRequestForEmptyUrl(t *testing.T) {
	_, err := SendRequest("")
	if err == nil {
		t.Errorf("Generate error on empty URL, got: nil")
	}
}

func TestSendRequestForInvalidUrl(t *testing.T) {
	_, err := SendRequest("some random garbage content")
	if err == nil {
		t.Errorf("Generate error on empty URL, got: nil")
	}
}

func TestSendRequestForValidUrl(t *testing.T) {
	_, err := SendRequest("http://www.google.com")
	if err != nil {
		t.Errorf("Should not get any error, got: %s", err)
	}
}

func TestProcessForInvalidUrl(t *testing.T) {
	err := Process("some garbage value")
	if err == nil {
		t.Errorf("Should return error, got nil")
	}
}

func TestProcessForValidUrl(t *testing.T) {
	err := Process("http://www.google.com")
	if err != nil {
		t.Errorf("Should not get any error, got %s", err)
	}
}

func TestConcurentExecutionForInValidUrls(t *testing.T) {
	tasks := make([]*Task, 0)
	tasks = append(tasks, NewTask(func() error { return Process("some random garbage value") }))
	tasks = append(tasks, NewTask(func() error { return Process("some more random garbage value") }))
	workerPool := NewWorkerPool(tasks, 2)
	workerPool.Run()
	for _, task := range workerPool.tasks {
		if task.err == nil {
			t.Errorf("Should return an error, got nil")
		}
	}
}

func TestConcurentExecutionForValidUrls(t *testing.T) {
	tasks := make([]*Task, 0)
	tasks = append(tasks, NewTask(func() error { return Process("http://www.google.com") }))
	tasks = append(tasks, NewTask(func() error { return Process("http://www.yahoo.com") }))
	workerPool := NewWorkerPool(tasks, 2)
	workerPool.Run()
	for _, task := range workerPool.tasks {
		if task.err != nil {
			t.Errorf("Should not get any error, got %s", task.err)
		}
	}
}

func TestConcurentExecutionForValidAndInvalidUrls(t *testing.T) {
	tasks := make([]*Task, 0)
	tasks = append(tasks, NewTask(func() error { return Process("http://www.google.com") }))
	tasks = append(tasks, NewTask(func() error { return Process("some random garbage") }))
	workerPool := NewWorkerPool(tasks, 2)
	workerPool.Run()
	errorCount := 0
	for _, task := range workerPool.tasks {
		if task.err != nil {
			errorCount++
		}
	}
	if errorCount > 1 {
		t.Errorf("Should only get one error, got %d", errorCount)
	}
}
