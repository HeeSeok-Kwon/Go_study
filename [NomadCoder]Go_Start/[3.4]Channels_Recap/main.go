package main

import (
	"fmt"
	"time"
)

func main() {
	// c := make(chan string)
	// people := [2]string{"nico", "flynn"}
	// for _, person := range people {
	// 	// result := go isSexy(person) // you can't do that
	// 	go isSexy(person, c)
	// }
	// time.Sleep(time.Second * 10)

	// result := <-c
	// fmt.Println(result)
	// fmt.Println("Waiting for messages")
	// resultOne := <-c
	// resultTwo := <-c
	// resultThree := <-c
	// fmt.Println("Received this message :", resultOne)
	// fmt.Println("Received this message :", resultTwo)
	// fmt.Println("Received this message :", resultThree)

	// fmt.Println("Waiting for messages")
	// fmt.Println("Received this message :", <-c) // blocking operation
	// fmt.Println("Received this message :", <-c)
	// fmt.Println(<-c) // deadlock -> Goroutines이 끝났기 때문

	// 원소가 추가된 경우 goroutines 처리 방법
	c := make(chan string)
	people := [5]string{"nico", "flynn", "dal", "japanguy", "larry"}
	for _, person := range people {
		go isSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}

}

// main이 끝나면 Goroutines도 끝남
// func isSexy(person string) bool {
// 	time.Sleep(time.Second * 5)
// 	return true
// }

// Channel : Goroutines - Main func 간의 커뮤니케이션
func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 10)
	// fmt.Println(person)
	c <- person + " is sexy"
}
