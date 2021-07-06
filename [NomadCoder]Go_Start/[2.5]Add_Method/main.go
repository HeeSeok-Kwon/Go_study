package main

import (
	"Nomad_Go_Start/modules/mydict"
	"fmt"
)

// package Nomad_Go_Start/mydict is not in GOROOT
// Project file에서 go mod init
func main() {
	// dictionary := mydict.Dictionary{}
	// dictionary["hello"] = "hello"
	// fmt.Println(dictionary)

	// dictionary := mydict.Dictionary{"first": "First word"}
	// fmt.Println(dictionary["first"])
	// definition, err := dictionary.Search("second")
	// definition, err := dictionary.Search("first")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }

	dictionary := mydict.Dictionary{}
	word := "hello"
	definition := "Greeting"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}

	hello, _ := dictionary.Search(word)
	fmt.Println("found", word, "definition:", hello)

	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}

}
