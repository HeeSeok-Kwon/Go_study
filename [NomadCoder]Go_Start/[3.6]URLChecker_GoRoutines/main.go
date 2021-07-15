package main

import (
	"errors"
	"fmt"
	"net/http"
)

type result struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request Failed")

func main() {
	results := make(map[string]string) // I don't use it --> error
	c := make(chan result)
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
		go hitURL(url, c)
	}
	// wait for message code needed
}

// send only
func hitURL(url string, c chan<- result) {
	fmt.Println("Checking:", url)
	// fmt.Println(<-c) // 내가 원하면 channel로부터 받을 수도 있다.
	// c <- result{} // channel로 보내고
	resp, err := http.Get(url)
	// if err != nil || resp.StatusCode >= 400 {
	// 	c <- result{url: url, status: "FAILED"}
	// } else {
	// 	c <- result{url: url, status: "OK"}
	// }
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- result{url: url, status: status}

}
