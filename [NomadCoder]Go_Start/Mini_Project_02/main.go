package main

import (
	"Nomad_Go_Start/modules/score"
	"fmt"
)

func main() {
	myScore := score.Score{}
	// Search test
	// s, err := myScore.Search("English")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("The Score is ", s)
	// }

	// Record test
	// err := myScore.Record("English", 90)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Record is successful")
	// }

	// Update test
	// myScore.Record("English", 90)
	// fmt.Println("The original score is", myScore["English"])
	// err := myScore.Update("English", 100)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Update is successful")
	// 	fmt.Println("The updated score is", myScore["English"])
	// }

	// Delete test
	// myScore.Record("English", 90)
	fmt.Println(myScore)
	err := myScore.Delete("English")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Delete is successful")
	}
	fmt.Println(myScore)
}
