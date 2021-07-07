package main

import (
	"Nomad_Go_Start/modules/mydict"
	"fmt"
)

// package Nomad_Go_Start/mydict is not in GOROOT
// Project file에서 go mod init --> 여기서 Nomad_Go_Start가 프로젝트 파일
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

	// dictionary := mydict.Dictionary{}
	// word := "hello"
	// definition := "Greeting"
	// err := dictionary.Add(word, definition)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// hello, _ := dictionary.Search(word)
	// fmt.Println("found", word, "definition:", hello)

	// err2 := dictionary.Add(word, definition)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }

	// dictionary := mydict.Dictionary{}
	// baseWord := "hello"
	// dictionary.Add(baseWord, "First")
	// // err := dictionary.Update(baseWord, "Second")
	// err := dictionary.Update("baseWord", "Second")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// word, _ := dictionary.Search(baseWord)
	// fmt.Println(word)

	// dictionary := mydict.Dictionary{}
	// baseWord := "hello"
	// dictionary.Add(baseWord, "First")
	// dictionary.Search(baseWord)
	// dictionary.Delete(baseWord)
	// word, err := dictionary.Search(baseWord)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(word)
	// }

	// My Delete Test
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Delete(baseWord)
	word, _ := dictionary.Search(baseWord)
	err := dictionary.Delete(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
}
