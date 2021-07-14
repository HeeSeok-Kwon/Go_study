package main

import (
	"fmt"
	"time"
)

func main() {
	// 원소가 추가된 경우 goroutines 처리 방법
	c := make(chan string) // 어떤 type을 받을 건지 지정
	people := [5]string{"nico", "flynn", "dal", "japanguy", "larry"}
	for _, person := range people {
		go isSexy(person, c)
	}
	// for i := 0; i < len(people); i++ {
	// 	fmt.Println("waiting for ", i, " ")
	// 	fmt.Println(<-c) // receive one msg
	// }
	fmt.Println(<-c) // 1 msg만 받고 나머지는 lost

}

// Channel : Goroutines - Main func 간의 커뮤니케이션
func isSexy(person string, c chan string) { // 어떤 type을 보낼 건지 지정
	time.Sleep(time.Second * 10)
	// fmt.Println(person)
	c <- person + " is sexy"
}

// 1. main function 끝나면 goroutines die
// 2. 어떤 type을 주고받을 건지 지정해야 함
// 3. channel <- : 메시지 보내는 방법
