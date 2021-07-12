package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request Failed")

func main() {
	// 초기화하지 않은 map 사용시 panic 발생
	// var results map[string]string
	// results["hello"] = "Hello"

	// 초기화 방법 1
	// var results = map[string]string{}
	// results["hello"] = "Hello"

	// 초기화 방법 2 --> otherwise map := nil
	var results = make(map[string]string)
	// results["hello"] = "Hello"

	urls := []string{
		"https://www.aribnb.com",
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.reddit.com",
		"https://soundcloud.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"https://acadmy.nomadcoders.co/",
	}

	for _, url := range urls {
		// fmt.Println(url)
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	// fmt.Println(results)
	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	fmt.Println(err, resp.StatusCode)
	return nil
}
