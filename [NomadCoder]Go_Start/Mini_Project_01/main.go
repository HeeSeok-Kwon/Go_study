package main

import (
	"Nomad_Go_Start/modules/mydict/school"
	"fmt"
)

func main() {
	// use public struct
	// univ := school.School{Name: "Pusan", Student: 150}
	// fmt.Println(univ.Name, univ.Student)
	// univ.Name = "Pukyong"
	// fmt.Println(univ.Name, univ.Student)

	// use Establish funcion
	univ := school.Establish("Pusan")
	fmt.Println(univ)

	// use Admission function
	fmt.Println(univ.Students())
	err := univ.Admission(301)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Admission completed")
	}
	fmt.Println("Total students :", univ.Students())
	// use Graduated function
	univ.Graduated(30)
	fmt.Println("Total students :", univ.Students())

	// use Rename function
	univ.Rename("Pukyong")
	// use Name function
	fmt.Println(univ.Name())

	fmt.Println(univ)

}
