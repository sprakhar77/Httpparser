package main

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"
)

func ComputeMD5Hash(data string) string {
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

func SendRequest(url string) ([]byte, error) {
	response, error := http.Get(url)

	if error != nil {
		return []byte{}, error
	}

	defer response.Body.Close()
	body, error := ioutil.ReadAll(response.Body)

	if error != nil {
		return []byte{}, error
	}

	return body, nil
}

func Process(url string) error {
	responseBody, err := SendRequest(url)
	if err != nil {
		log.Println("Could not parse resonse for url:", url, "error:", err)
		return err
	}

	md5hash := ComputeMD5Hash(string(responseBody))
	log.Println(url, md5hash)
	return nil
}
