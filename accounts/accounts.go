package accounts

// Account private struct
type account struct {
	owner   string
	balance int
}

// NewAccoutn create Account
func NewAccount(owner string) *account {
	account := account{owner: owner, balance: 0}
	return &account
}
