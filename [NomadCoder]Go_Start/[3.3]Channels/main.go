package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		// result := go isSexy(person) // you can't do that
		go isSexy(person, c)
	}
	// time.Sleep(time.Second * 10)

	// result := <-c
	// fmt.Println(result)
	fmt.Println(<-c)
	fmt.Println(<-c)
	// fmt.Println(<-c) // deadlock -> Goroutines이 끝났기 때문
}

// main이 끝나면 Goroutines도 끝남
// func isSexy(person string) bool {
// 	time.Sleep(time.Second * 5)
// 	return true
// }

// Channel : Goroutines - Main func 간의 커뮤니케이션
func isSexy(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	c <- true
}
