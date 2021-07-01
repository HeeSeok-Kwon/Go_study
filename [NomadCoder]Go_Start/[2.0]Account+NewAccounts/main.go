package main

import (
	"accounts"
	"fmt"
)

func main() {
	// use public struct
	// account := banking.Account{Owner: "nicolas", Balance: 1000}
	// account.Owner = "pepe"

	// use NewAccount function
	account := accounts.NewAccount("nico")
	// account.owner = "pepe" // 바꿀 수 없다.
	// account.balance = 1000 // 바꿀 수 없다.
	fmt.Println(account)
}
