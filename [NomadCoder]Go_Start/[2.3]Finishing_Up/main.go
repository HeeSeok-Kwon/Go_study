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
	account.Deposit(10)
	// fmt.Println(account.Balance())
	// err := account.Withdraw(20)
	// if err != nil {
	// 	// log.Fatalln(err)
	// 	fmt.Println(err)
	// }
	// fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account)
}
