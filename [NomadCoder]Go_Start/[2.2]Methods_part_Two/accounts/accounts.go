package accounts

import "errors"

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

var errNoMoney = errors.New("Can't withdraw")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// method --> struct의 첫 글자를 따서 소문자로 해야 함
// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	// fmt.Println("Gonna deposit", amount)
	a.balance += amount
	// fmt.Println("New balance", a.balance)
}

// Balance of your account
func (a Account) Balance() int {
	// fmt.Println(a.balance)
	return a.balance
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		// return errors.New("Can't withdraw you are poor")
		return errNoMoney
	}
	a.balance -= amount
	return nil
}
