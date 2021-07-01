package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	// nico := map[string]string{"nico": "string"}
	// fmt.Println(nico)

	favFood := []string{"kimchi", "ramen"}
	// nico := person{"name", 18, favFood} // not clear code
	nico := person{name: "nico", age: 18, favFood: favFood} // more clear
	fmt.Println(nico)
	fmt.Println(nico.name)
}
