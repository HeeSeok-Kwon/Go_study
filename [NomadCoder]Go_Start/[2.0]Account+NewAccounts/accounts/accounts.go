package accounts

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
