package main 

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	fmt.Println(multiply(2, 3))

	totalLength, upperName := lenAndUpper("nico")
	fmt.Println(totalLength, upperName)

	buf, _ := lenAndUpper("nico")
	fmt.Println(buf)
	
	repeatMe("nico", "lynn", "dal", "marl", "flynn")
}
