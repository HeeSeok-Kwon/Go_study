package main

import (
	"fmt"
	"time"
)

func main() {
	// top - down
	// sexyCount("nico")
	// sexyCount("flynn")

	// Goroutines
	// go sexyCount("nico")
	// sexyCount("flynn")

	// Goroutines은 main func이 실행될 때만 유효, 존재
	go sexyCount("nico")
	go sexyCount("flynn")
	time.Sleep(time.Second * 5)
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
