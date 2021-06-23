package main

import "fmt"

func superAdd(numbers ...int) int {
	// for index, number := range numbers {
	// 	fmt.Println(index, number)
	// }
	// return 1

	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }
	// return 1

	total := 0
	for _, number := range numbers {
		total = total + number
		// total += number
	}

	return total
}

func main() {
	// superAdd(1, 2, 3, 4, 5, 6)
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}
