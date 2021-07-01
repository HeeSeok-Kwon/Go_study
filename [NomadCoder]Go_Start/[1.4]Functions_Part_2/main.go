package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int, uppercase string) {
	//defer -- function이 return하면 실행됨
	defer fmt.Println("I'm done")
	//naked return
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
	//return length, uppercase
}

func main() {
	totalLength, up := lenAndUpper("nico")
	fmt.Println(totalLength, up)
}
