package accounts

import "fmt"

// Account struct - public
// type Account struct {
// 	Owner   string
// 	Balance int
// }

// Account struct - private
type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// method --> struct의 첫 글자를 따서 소문자로 해야 함
// Deposit x amount on your account
func (a Account) Deposit(amount int) {
	fmt.Println("Gonna deposit", amount)
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}
